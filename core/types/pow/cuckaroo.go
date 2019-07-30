package pow

import (
	"errors"
	"github.com/HalalChain/qitmeer-lib/common/hash"
	"github.com/HalalChain/qitmeer-lib/crypto/cuckoo"
	"github.com/HalalChain/qitmeer-lib/log"
)

const (
	EDGE_BITS_START = 4
	EDGE_BITS_END = 8
	CIRCLE_NONCE_START = 12
	CIRCLE_NONCE_END = 180
)

type Cuckaroo struct {
	Pow
}

func (this *Cuckaroo) Verify(h hash.Hash,targetDiff uint64) error{
	nonces := this.GetCircleNonces()
	err := cuckoo.VerifyCuckaroo(h[:],nonces[:])
	if err != nil{
		log.Debug("Verify Error!",err)
		return err
	}
	if CalcCuckooDiff(this.GetScale(),this.GetBlockHash([]byte{})) < targetDiff{
		return errors.New("difficulty is too easy!")
	}
	return nil
}

func (this *Cuckaroo)GetBlockHash (data []byte) hash.Hash {
	circlNonces := [42]uint64{}
	nonces := this.GetCircleNonces()
	for i:=0;i<len(nonces);i++{
		circlNonces[i] = uint64(nonces[i])
	}
	return CuckooHash(circlNonces[:],int(this.GetEdgeBits()))
}

func (this *Cuckaroo) GetNonce () uint64 {
	return this.Nonce
}

func (this *Cuckaroo) GetPowType () PowType {
	return CUCKAROO
}