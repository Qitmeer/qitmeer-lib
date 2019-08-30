package pow

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/Qitmeer/qitmeer-lib/common/hash"
	"math/big"
)

// proof data length 188
const POW_LENGTH = 188
//except pow type 4bytes and nonce 8 bytes 176 bytes
const PROOFDATA_LENGTH = 176
type PowType int
type PowBytes [POW_LENGTH]byte

const (
	//pow type enum
	BLAKE2BD PowType = 0
	CUCKAROO PowType = 1
	CUCKATOO PowType = 2
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
	// PowLimit defines the highest allowed proof of work value for a block
	// as a uint256.
	Blake2bdPowLimit *big.Int
	// PowLimitBits defines the highest allowed proof of work value for a
	// block in compact form.
	// highest value is mean min difficulty
	Blake2bdPowLimitBits uint32

	// cuckoo difficulty calc params  min difficulty
	CuckarooMinDifficulty uint32
	CuckatooMinDifficulty uint32
	// solotion difficulty diff = (scale * 1<<64) u128 / cuckoohash as u128
	CuckarooDiffScale uint64
	CuckatooDiffScale uint64

	//percent of every pow sum of them must be 100
	CuckarooPercent int
	CuckatooPercent int
	Blake2bDPercent int
}

type IPow interface {
	// verify result difficulty
	Verify(headerWithoutProofData []byte,targetDiff uint64,powConfig *PowConfig) error
	//set header nonce
	SetNonce(nonce uint64)
	//calc next diff
	GetNextDiffBig(weightedSumDiv *big.Int,oldDiffBig *big.Int,currentPowPercent *big.Int,param *PowConfig) *big.Int
	GetNonce() uint64
	GetPowType() PowType
	//set pow type
	SetPowType(powType PowType)
	GetProofData() string
	//set proof data
	SetProofData([]byte)
	GetBlockHash(data []byte) hash.Hash
	Bytes() PowBytes
	//if cur_reduce_diff > 0 compare cur_reduce_diff with powLimitBits or minDiff ，the cur_reduce_diff should less than powLimitBits , and should more than min diff
	//if cur_reduce_diff <=0 return powLimit or min diff
	GetSafeDiff(param *PowConfig,cur_reduce_diff uint64) uint64
	PowPercent(param *PowConfig) *big.Int
}


type Pow struct {
	PowType PowType //header pow type 4 bytes
	Nonce uint64 //header nonce 8 bytes
	ProofData ProofDataType // 4 edge_bits + 4 diff scale + 168  bytes circle length total 176 bytes
}


func (this *Pow)Bytes() PowBytes {
	r := [POW_LENGTH]byte{}
	//write pow type 4 bytes
	n := make([]byte,4)
	binary.LittleEndian.PutUint32(n,uint32(this.PowType))
	copy(r[0:4],n)
	//write nonce 8 bytes
	n = make([]byte,8)
	binary.LittleEndian.PutUint64(n,this.Nonce)
	copy(r[4:12],n)
	//write ProofData 176 bytes
	copy(r[12:],this.ProofData[:])
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
	this.PowType = powType
}

func (this *Pow) GetPowType () PowType {
	return this.PowType
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
	copy(this.ProofData[0:l],data[:])
}

//check pow is available
func (this *Pow) CheckAvailable (powPercent *big.Int) bool {
	return powPercent.Cmp(big.NewInt(0)) <=0
}