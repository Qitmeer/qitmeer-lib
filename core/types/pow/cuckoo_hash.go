package pow

import (
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/common/util"
	"sort"
)

func CuckooHash(nonces []uint64,nonce_bits int) hash.Hash {
	sort.Slice(nonces, func(i, j int) bool {
		return nonces[i] < nonces[j]
	})
	bitvec,_ := util.New(nonce_bits*42)
	for i:=41;i>=0;i--{
		n := i
		nonce := nonces[i]
		for bit:= 0;bit < nonce_bits;bit++{
			if nonce & (1 << uint(bit)) != 0 {
				bitvec.SetBitAt(n * nonce_bits + bit)
			}
		}
	}
	return hash.HashH(bitvec.Bytes())
}
