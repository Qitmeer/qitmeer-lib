package pow

import (
	"encoding/binary"
	"github.com/HalalChain/qitmeer-lib/common/hash"
)

type Cuckatoo struct {
	Pow
}

func (this *Cuckatoo) Verify(data []byte) bool{
	return true
}

func (this *Cuckatoo) CalcScale () int {
	return 1
}
func (this *Cuckatoo)GetBlockHash (data []byte) hash.Hash {
	circlNonces := []uint64{}
	nonces := this.GetCircleNonces()
	for i:=0;i<len(nonces);i++{
		circlNonces[i] = uint64(nonces[i])
	}
	return CuckooHash(circlNonces,int(this.GetEdgeBits()))
}

func (this *Cuckatoo) GetNonce () uint64 {
	return this.Nonce
}

func (this *Cuckatoo) GetPowType () PowType {
	return CUCKATOO
}

func (this *Cuckatoo) GetEdgeBits () uint32 {
	return binary.LittleEndian.Uint32(this.ProofData[EDGE_BITS_START:EDGE_BITS_END])
}

func (this *Cuckatoo) GetCircleNonces () (nonces [42]uint32) {
	nonces = [42]uint32{}
	j := 0
	for i :=CIRCLE_NONCE_START;i<CIRCLE_NONCE_END;i+=4{
		nonceBytes := this.ProofData[i:i+4]
		nonces[j] = binary.LittleEndian.Uint32(nonceBytes)
		j++
	}
	return
}