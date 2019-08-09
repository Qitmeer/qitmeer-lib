package pow

import (
	"errors"
	"fmt"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/crypto/cuckoo"
	"github.com/HalalChain/qitmeer-lib/log"
	"math/big"
)

type Cuckatoo struct {
	Cuckoo
}

const MIN_CUCKATOOEDGEBITS = 29

func (this *Cuckatoo) Verify(headerWithoutProofData []byte,targetDiff uint64) error{
	h := hash.HashH(headerWithoutProofData)
	nonces := this.GetCircleNonces()
	edgeBits := this.GetEdgeBits()
	if edgeBits < MIN_CUCKATOOEDGEBITS{
		return fmt.Errorf("edge bits:%d is too short!less than %d",edgeBits,MIN_CUCKATOOEDGEBITS)
	}
	err := cuckoo.VerifyCuckatoo(h[:],nonces[:],uint(edgeBits))
	if err != nil{
		log.Debug("Verify Error!",err)
		return err
	}
	if this.CalcCuckooDiff(this.GetScale(),this.GetBlockHash([]byte{})) < targetDiff{
		return errors.New("difficulty is too easy!")
	}
	return nil
}

func (this *Cuckatoo) GetNextDiffBig(weightedSumDiv *big.Int,oldDiffBig *big.Int,currentPowPercent *big.Int,param *PowConfig) *big.Int{
	nextDiffBig := oldDiffBig.Div(oldDiffBig, weightedSumDiv)
	targetPercent := this.PowPercent(param)
	if currentPowPercent.Cmp(targetPercent) > 0{
		currentPowPercent.Div(currentPowPercent,targetPercent)
		nextDiffBig.Mul(nextDiffBig,currentPowPercent)
	}
	return nextDiffBig
}
func (this *Cuckatoo) PowPercent(param *PowConfig) *big.Int{
	targetPercent := big.NewInt(int64(param.CuckatooPercent))
	targetPercent.Lsh(targetPercent,32)
	return targetPercent
}

func (this *Cuckatoo) GetMinDiff(param *PowConfig) uint64{
	return uint64(param.CuckatooPowLimitBits)
}