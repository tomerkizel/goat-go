package goat

import (
	"fmt"
	"sort"

	"github.com/tomerkizel/goat-go/utils"
)

type PArray struct {
	elemtype   any
	arrayValue []any
}

func EmptyPArray(elemtype any) *PArray {
	self := PArray{elemtype, nil}
	return &self
}

func (p *PArray) Push(elem any) (*PArray, error) {
	e := utils.CheckType(elem, p.elemtype)
	if e != nil {
		return nil, e
	}
	pn := EmptyPArray(p.elemtype)
	pn.arrayValue = make([]any, len(p.arrayValue)+1)
	final := copy(pn.arrayValue, p.arrayValue)
	pn.arrayValue[final] = elem
	return pn, nil
}

func (p *PArray) Merge(q *PArray) (*PArray, error) {
	err := utils.CheckType(p.elemtype, q.elemtype)
	if err != nil {
		return nil, err
	}
	pn := EmptyPArray(p.elemtype)
	pn.arrayValue = make([]any, len(p.arrayValue)+len(q.arrayValue))
	final := copy(pn.arrayValue, p.arrayValue)
	for _, v := range q.GetArray() {
		pn.arrayValue[final] = v
		final++
	}
	return pn, nil
}

func (p *PArray) Set(index int, elem any) (*PArray, error) {
	if index < 0 || index > len(p.arrayValue) {
		return nil, fmt.Errorf("index %v out of range for length %v", index, len(p.arrayValue))
	}
	e := utils.CheckType(elem, p.elemtype)
	if e != nil {
		return nil, e
	}
	pn := EmptyPArray(p.elemtype)
	pn.arrayValue = make([]any, len(p.arrayValue))
	copy(pn.arrayValue, p.arrayValue)
	pn.arrayValue[index] = elem
	return pn, nil
}

func (p *PArray) Get(index int) (any, error) {
	if index > len(p.arrayValue) || index < 0 {
		return nil, fmt.Errorf("index %v out of range for length %v", index, len(p.arrayValue))
	}
	return p.arrayValue[index], nil
}

func (p *PArray) Delete(index int) (*PArray, any, error) {
	if index > len(p.arrayValue) {
		return nil, nil, fmt.Errorf("index %v out of range for length %v", index, len(p.arrayValue))
	}
	pn := EmptyPArray(p.elemtype)
	pn.arrayValue = make([]any, len(p.arrayValue)-1)
	count := 0
	for i := range p.arrayValue {
		if i == index {
			continue
		}
		pn.arrayValue[count] = p.arrayValue[i]
		count++
	}
	return pn, p.arrayValue[index], nil
}

func (p *PArray) Pop() (*PArray, any, error) {
	if len(p.arrayValue) == 0 {
		return nil, nil, fmt.Errorf("can't pop an empty array")
	}
	return p.Delete(len(p.arrayValue) - 1)
}

func (p *PArray) Sort(fn func(x, y any) bool) (*PArray, error) {
	pn := EmptyPArray(p.elemtype)
	pn.arrayValue = make([]any, len(p.arrayValue))
	copy(pn.arrayValue, p.arrayValue)
	sort.Slice(pn.arrayValue, func(i, j int) bool {
		return fn(pn.arrayValue[i], pn.arrayValue[j])
	})
	return pn, nil
}

func (p *PArray) GetArray() []any {
	pn := make([]any, len(p.arrayValue))
	copy(pn, p.arrayValue)
	return pn
}
