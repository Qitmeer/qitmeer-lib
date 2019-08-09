package pow

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"math/big"
)

// proof data length 204
const POW_LENGTH = 208
const PROOFDATA_LENGTH = 200
type PowType int
type PowBytes [POW_LENGTH]byte

const (
	BLAKE2BD PowType = 0
	CUCKAROO PowType = 1
	CUCKATOO PowType = 2
	POW_TYPE_START = 0
	POW_TYPE_END = 4
)

var PowMapString = map[PowType]interface{}{
	BLAKE2BD : "blake2bd",
	CUCKAROO : "cuckaroo",
	CUCKATOO : "cuckatoo",
}

type ProofDataType [PROOFDATA_LENGTH]byte

func (this *ProofDataType) String() string{
	return hex.EncodeToString(this[:])
}

func (this *ProofDataType) Bytes() []byte{
	return this[:]
}

type PowConfig struct {
	// PowLimitBits defines the highest allowed proof of work value for a
	// block in compact form.
	PowLimitBits uint32

	// Cuckoo PowLimitBits defines the highest allowed proof of work value for a
	// block in compact form.
	CuckarooPowLimitBits uint32
	CuckatooPowLimitBits uint32
	CuckarooScale uint64
	CuckatooScale uint64
	//percent of pow
	CuckarooPercent int
	CuckatooPercent int
	Blake2bDPercent int
}

type IPow interface {
	Verify(headerWithoutProofData []byte,targetDiff uint64) error
	SetNonce(nonce uint64)
	GetNextDiffBig(weightedSumDiv *big.Int,oldDiffBig *big.Int,currentPowPercent *big.Int,param *PowConfig) *big.Int
	GetNonce() uint64
	GetPowType() PowType
	SetPowType(powType PowType)
	GetProofData() string
	SetProofData([]byte)
	GetBlockHash(data []byte) hash.Hash
	Bytes() PowBytes
	GetMinDiff(param *PowConfig) uint64
	PowPercent(param *PowConfig) *big.Int
}


type Pow struct {
	Nonce uint64 //header nonce 8 bytes
	ProofData ProofDataType // 4 powType + 4 edge_bits + 4 diff scale + 168  bytes circle length + other 20 bytes ... may other new pow proof data struct ,but the first 4 bytes must be pwo type
}

func (this *Pow)Bytes() PowBytes {
	nonceLen := POW_LENGTH - PROOFDATA_LENGTH
	r := [POW_LENGTH]byte{}
	n := make([]byte,nonceLen)
	binary.LittleEndian.PutUint64(n,this.Nonce)
	copy(r[0:nonceLen],n)
	copy(r[nonceLen:],this.ProofData[:])
	return PowBytes(r)
}

//get pow instance
func GetInstance (powType PowType,nonce uint64,proofData []byte) IPow {
	var instance IPow
	switch powType {
	case BLAKE2BD:
		instance = &Blake2bd{}
	case CUCKAROO:
		instance = &Cuckaroo{}
	case CUCKATOO:
		instance = &Cuckatoo{}
	default:
		instance = &Blake2bd{}
	}
	instance.SetPowType(powType)
	instance.SetNonce(nonce)
	instance.SetProofData(proofData)
	return instance
}

func (this *Pow) SetPowType (powType PowType) {
	binary.LittleEndian.PutUint32(this.ProofData[POW_TYPE_START:POW_TYPE_END],uint32(powType))
}

func (this *Pow) GetPowType () PowType {
	return PowType(binary.LittleEndian.Uint32(this.ProofData[POW_TYPE_START:POW_TYPE_END]))
}

func (this *Pow) GetNonce () uint64 {
	return this.Nonce
}

func (this *Pow) SetNonce (nonce uint64) {
	this.Nonce = nonce
}

func (this *Pow) GetProofData () string {
	return this.ProofData.String()
}

//set proof data except pow type
func (this *Pow) SetProofData (data []byte) {
	l := len(data)
	copy(this.ProofData[POW_TYPE_END:l+POW_TYPE_END],data[:])
}