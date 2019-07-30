package pow

import (
	"encoding/binary"
	"github.com/HalalChain/qitmeer-lib/common/hash"
)

// proof data length 204
const PROOFDATA_LENGTH = 200
type PowType int
type PowBytes [208]byte

const (
	BLAKE2BD PowType = 0
	CUCKAROO PowType = 1
	CUCKATOO PowType = 2

	POW_TYPE_START = 128
	POW_TYPE_END = 132
	POW_START = 120
	POW_END = 328

	DIFF_START = 100
	DIFF_END = 104
)

type IPow interface {
	Verify(h hash.Hash,targetDiff uint64) error
	GetNonce() uint64
	GetPowType() PowType
	GetBlockHash(data []byte) hash.Hash
	Bytes() PowBytes
}


type Pow struct {
	Nonce uint64 //header nonce
	ProofData [PROOFDATA_LENGTH]byte // 4 powType + 4 edge_bits + 200 bytes circle length ... may other new pow proof data struct ,but the first 4 bytes must be pwo type
}

func (this *Pow)Bytes() PowBytes {
	r := [208]byte{}
	n := make([]byte,8)
	binary.LittleEndian.PutUint64(n,this.Nonce)
	copy(r[0:8],n)
	copy(r[8:208],this.ProofData[:])
	return PowBytes(r)
}

func GetInstance (powType PowType) IPow {
	switch powType {
	case BLAKE2BD:
		instance := &Blake2bd{}
		binary.LittleEndian.PutUint32(instance.ProofData[:4],uint32(powType))
		return instance
	case CUCKAROO:
		instance := &Cuckaroo{}
		binary.LittleEndian.PutUint32(instance.ProofData[:4],uint32(powType))
		return instance
	case CUCKATOO:
		instance := &Cuckatoo{}
		binary.LittleEndian.PutUint32(instance.ProofData[:4],uint32(powType))
		return instance
	default:
		instance := &Blake2bd{}
		binary.LittleEndian.PutUint32(instance.ProofData[:4],uint32(powType))
		return instance
	}
}

func (this *Pow) GetCircleNonces () (nonces [42]uint32) {
	nonces = [42]uint32{}
	j := 0
	for i :=CIRCLE_NONCE_START;i<CIRCLE_NONCE_END;i+=4{
		nonceBytes := this.ProofData[i:i+4]
		nonces[j] = binary.LittleEndian.Uint32(nonceBytes)
		j++
	}
	return
}

func (this *Pow) GetEdgeBits () uint32 {
	return binary.LittleEndian.Uint32(this.ProofData[EDGE_BITS_START:EDGE_BITS_END])
}