package pow

import (
	"errors"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/crypto/cuckoo"
	"github.com/HalalChain/qitmeer-lib/log"
)

type Cuckatoo struct {
	Pow
}

func (this *Cuckatoo) Verify(h hash.Hash,targetDiff uint64) error{
	nonces := this.GetCircleNonces()
	err := cuckoo.VerifyCuckatoo(h[:],nonces[:])
	if err != nil{
		log.Error("Verify Error!",err)
		return err
	}
	if CalcCuckooDiff(this.CalcScale(),this.GetBlockHash([]byte{})) < targetDiff{
		return errors.New("difficulty is too easy!")
	}
	return nil
}

func (this *Cuckatoo) CalcScale () int64 {
	return 1856
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
