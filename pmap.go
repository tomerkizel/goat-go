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

func (p *PMap) AddOne(key, value any) (*PMap, error) {
	pn := PMap{}
	pn.typekey = p.typekey
	pn.typevalue = p.typevalue
	pn.mapValue = make(map[any]any)
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
func (p *PMap) AddBatch(keyvalue map[any]any) (*PMap, error) {
	pn := PMap{}
	pn.typekey = p.typekey
	pn.typevalue = p.typevalue
	pn.mapValue = make(map[any]any)
	for k, v := range p.mapValue {
		pn.mapValue[k] = v
	}
	for key, value := range keyvalue {
		err := utils.CheckType(key, pn.typekey)
		if err != nil {
			return nil, err
		}
		err = utils.CheckType(value, pn.typevalue)
		if err != nil {
			return nil, err
		}

		pn.mapValue[key] = value
	}
	return &pn, nil
}

func (p *PMap) Read(key any) (any, error) {
	e := utils.CheckType(key, p.typekey)
	if e != nil {
		return nil, e
	}
	return p.mapValue[key], nil
}

func (p *PMap) Delete(key any) (*PMap, error) {
	e := utils.CheckType(key, p.typekey)
	if e != nil {
		return nil, e
	}
	pn := PMap{}
	pn.typekey = p.typekey
	pn.typevalue = p.typevalue
	pn.mapValue = make(map[any]any)
	for k, v := range p.mapValue {
		if k == key {
			continue
		}
		pn.mapValue[k] = v
	}
	return &pn, nil
}
