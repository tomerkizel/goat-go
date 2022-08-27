package goat

import (
	"fmt"

	"github.com/tomerkizel/goat-go/utils"
)

type PArray struct {
	elemtype   any
	arrayValue []any
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
	pn.arrayValue = make([]any, len(p.arrayValue)+1)
	final := copy(pn.arrayValue, p.arrayValue)
	pn.arrayValue[final] = elem
	return &pn, nil
}

func (p *PArray) AddBatch(elems []any) (*PArray, error) {
	pn := PArray{}
	pn.elemtype = p.elemtype
	pn.arrayValue = make([]any, len(p.arrayValue)+len(elems))
	final := copy(pn.arrayValue, p.arrayValue)
	for i := range elems {
		err := utils.CheckType(elems[i], pn.elemtype)
		if err != nil {
			return nil, err
		}
		pn.arrayValue[final] = elems[i]
		final++
	}
	return &pn, nil
}

func (p *PArray) Read(index int) (any, error) {
	if index > len(p.arrayValue) {
		return nil, fmt.Errorf("index %v out of range for length %v", index, len(p.arrayValue))
	}
	return p.arrayValue[index], nil
}

func (p *PArray) Delete(index int) (*PArray, error) {
	if index > len(p.arrayValue) {
		return nil, fmt.Errorf("index %v out of range for length %v", index, len(p.arrayValue))
	}
	pn := PArray{}
	pn.elemtype = p.elemtype
	pn.arrayValue = make([]any, len(p.arrayValue)-1)
	count := 0
	for i := range p.arrayValue {
		if i == index {
			continue
		}
		pn.arrayValue[count] = p.arrayValue[i]
		count++
	}
	return &pn, nil
}
