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
)

type IPow interface {
	Verify(data []byte) bool
	CalcScale() int
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
