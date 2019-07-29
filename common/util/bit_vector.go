package util

import "errors"

var errInvalidLength = errors.New("invalid length")

type BitVector struct {
	len int
	b   []byte
}

func New(l int) (bv *BitVector, err error) {
	return NewFromBytes(make([]byte, l/8+1), l)
}

func NewFromBytes(b []byte, l int) (bv *BitVector, err error) {
	if l <= 0 {
		return nil, errInvalidLength
	}
	if len(b)*8 < l {
		return nil, errInvalidLength
	}
	return &BitVector{
		len: l,
		b:   b,
	}, nil
}

func (bv *BitVector) Get(i int) bool {
	bi := i / 8
	return bv.b[bi]&(0x1<<uint(i%8)) != 0
}

func (bv *BitVector) Set(i int, v bool) {
	bi := i / 8
	cv := bv.Get(i)
	if cv != v {
		bv.b[bi] ^= 0x1 << uint8(i%8)
	}
}

func (bv *BitVector) SetBitAt(pos int) {
	bv.b[pos / 8] |= 1 << (uint(pos) % 8)
}

func (bv *BitVector) Bytes() []byte {
	return bv.b
}