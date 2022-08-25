package pmap

import (
	"github.com/tomerkizel/goat-go/utils"
)

type PerMap struct {
	typekey   any
	typevalue any
	mapValue  map[any]any
}

func Empty(keytype, valtype any) *PerMap {
	self := PerMap{keytype, valtype, make(map[any]any)}
	return &self
}

func (p *PerMap) Add(key, value any) (*PerMap, error) {
	pn := PerMap{}
	pn.typekey = p.typekey
	pn.typevalue = p.typevalue
	err := utils.CheckType(key, pn.typekey)
	if err != nil {
		return nil, err
	}
	err = utils.CheckType(value, pn.typevalue)
	if err != nil {
		return nil, err
	}
	for k, v := range p.mapValue {
		pn.mapValue[k] = v
	}
	pn.mapValue[key] = value
	return &pn, nil
}
