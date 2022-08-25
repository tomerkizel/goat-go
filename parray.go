package goat

import (
	"github.com/tomerkizel/goat-go/utils"
)

type PArray struct {
	elemtype any
	valueArr []any
}

func EmptyPArray(elemtype any) *PArray {
	self := PArray{elemtype, make([]any, 0)}
	return &self
}

func (p *PArray) AddOne(elem any) (*PArray, error) {
	e := utils.CheckType(elem, p.elemtype)
	if e != nil {
		return nil, e
	}
	pn := PArray{}
	pn.elemtype = p.elemtype
	pn.valueArr = make([]any, len(p.valueArr)+1)
	final := copy(pn.valueArr, p.valueArr)
	pn.valueArr[final] = elem
	return &pn, nil
}
