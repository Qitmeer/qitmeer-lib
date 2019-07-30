package pow

import (
	"errors"
	"fmt"
	"github.com/HalalChain/qitmeer-lib/common/hash"
)

type Blake2bd struct {
	Pow
}

func (this *Blake2bd) Verify(h hash.Hash,targetDiff uint64) error{
	target := CompactToBig(uint32(targetDiff))
	if target.Sign() <= 0 {
		str := fmt.Sprintf("block target difficulty of %064x is too "+
			"low", target)
		return errors.New(str)
	}

	// The target difficulty must be less than the maximum allowed.
	//if target.Cmp(powLimit) > 0 {
	//	str := fmt.Sprintf("block target difficulty of %064x is "+
	//		"higher than max of %064x", target, powLimit)
	//	return errors.New(str)
	//}
	hashNum := HashToBig(&h)
	if hashNum.Cmp(target) > 0 {
		str := fmt.Sprintf("block hash of %064x is higher than"+
			" expected max of %064x", hashNum, target)
		return errors.New(str)
	}
	return nil
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