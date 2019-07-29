package pow

import "github.com/HalalChain/qitmeer-lib/common/hash"

type Blake2bd struct {
	Pow
}

func (this *Blake2bd) Verify(data []byte) bool{
	return true
}

func (this *Blake2bd) CalcScale () int {
	return 1
}

func (this *Blake2bd) GetNonce () uint64 {
	return this.Nonce
}

func (this *Blake2bd) GetPowType () PowType {
	return BLAKE2BD
}

func (this *Blake2bd)GetBlockHash (data []byte) hash.Hash {
	return hash.DoubleHashH(data)
}