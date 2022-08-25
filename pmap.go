package goat

import (
	"github.com/tomerkizel/goat-go/utils"
)

type PMap struct {
	typekey   any
	typevalue any
	mapValue  map[any]any
}

func EmptyPMap(keytype, valtype any) *PMap {
	self := PMap{keytype, valtype, make(map[any]any)}
	return &self
}

func (p *PMap) Add(key, value any) (*PMap, error) {
	pn := PMap{}
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
