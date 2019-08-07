package pow

import (
	"errors"
	"fmt"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"math/big"
)

type Blake2bd struct {
	Pow
}

func (this *Blake2bd) Verify(headerWithoutProofData []byte,targetDiff uint64) error{
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
	h := hash.DoubleHashH(headerWithoutProofData)
	hashNum := HashToBig(&h)
	if hashNum.Cmp(target) > 0 {
		str := fmt.Sprintf("block hash of %064x is higher than"+
			" expected max of %064x", hashNum, target)
		return errors.New(str)
	}
	return nil
}

func (this *Blake2bd)GetBlockHash (data []byte) hash.Hash {
	return hash.DoubleHashH(data)
}

func (this *Blake2bd) GetNextDiffBig(weightedSumDiv *big.Int,oldDiffBig *big.Int,currentPowPercent *big.Int,param *PowConfig) *big.Int{
	nextDiffBig := weightedSumDiv.Mul(weightedSumDiv, oldDiffBig)
	targetPercent := this.PowPercent(param)
	if currentPowPercent.Cmp(targetPercent) > 0{
		currentPowPercent.Div(currentPowPercent,targetPercent)
		nextDiffBig.Div(nextDiffBig,currentPowPercent)
	}
	return nextDiffBig
}

func (this *Blake2bd) PowPercent(param *PowConfig) *big.Int{
	targetPercent := big.NewInt(int64(param.Blake2bDPercent))
	targetPercent.Lsh(targetPercent,32)
	return targetPercent
}

func (this *Blake2bd) GetMinDiff(param *PowConfig) uint64{
	return uint64(param.PowLimitBits)
}