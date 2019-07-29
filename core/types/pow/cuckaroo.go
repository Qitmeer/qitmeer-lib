package pow

import (
	"encoding/binary"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/crypto/cuckoo"
	"github.com/HalalChain/qitmeer-lib/log"
)

const (
	EDGE_BITS_START = 4
	EDGE_BITS_END = 8
	CIRCLE_NONCE_START = 8
	CIRCLE_NONCE_END = 176
)

type Cuckaroo struct {
	Pow
}

func (this *Cuckaroo) Verify(data []byte) bool{
	nonces := this.GetCircleNonces()
	err := cuckoo.VerifyCuckaroo(data[:],nonces[:])
	if err != nil{
		log.Error("Verify Error!",err)
		return false
	}
	return true
}

func (this *Cuckaroo)GetBlockHash (data []byte) hash.Hash {
	circlNonces := []uint64{}
	nonces := this.GetCircleNonces()
	for i:=0;i<len(nonces);i++{
		circlNonces[i] = uint64(nonces[i])
	}
	return CuckooHash(circlNonces,int(this.GetEdgeBits()))
}

func (this *Cuckaroo) CalcScale () int {
	return 1
}

func (this *Cuckaroo) GetNonce () uint64 {
	return this.Nonce
}

func (this *Cuckaroo) GetPowType () PowType {
	return CUCKAROO
}

func (this *Cuckaroo) GetEdgeBits () uint32 {
	return binary.LittleEndian.Uint32(this.ProofData[EDGE_BITS_START:EDGE_BITS_END])
}

func (this *Cuckaroo) GetCircleNonces () (nonces [42]uint32) {
	nonces = [42]uint32{}
	j := 0
	for i :=CIRCLE_NONCE_START;i<CIRCLE_NONCE_END;i+=4{
		nonceBytes := this.ProofData[i:i+4]
		nonces[j] = binary.LittleEndian.Uint32(nonceBytes)
		j++
	}
	return
}