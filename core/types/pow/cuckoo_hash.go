package pow

import (
	"encoding/binary"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/common/util"
	"math/big"
	"sort"
)

//calc cuckoo hash
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
	h := hash.HashH(bitvec.Bytes())
	util.ReverseBytes(h[:])
	return h
}

//calc cuckoo diff
func CalcCuckooDiff(scale int64,blockHash hash.Hash) uint64 {
	c := &big.Int{}
	util.ReverseBytes(blockHash[:])
	c.SetUint64(binary.BigEndian.Uint64(blockHash[:8]))
	a := big.NewInt(scale)
	d := big.NewInt(1)
	d.Lsh(d,64)
	a.Mul(a,d)
	e := a.Div(a,c)
	return e.Uint64()
}
