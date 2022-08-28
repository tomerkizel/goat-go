package goat

import (
	"github.com/tomerkizel/goat-go/utils"
)

type PMap struct {
	keyType   any
	valueType any
	mapValue  map[any]any
}

func EmptyPMap(keytype, valtype any) *PMap {
	self := PMap{keytype, valtype, make(map[any]any)}
	return &self
}

func (p *PMap) Set(key, value any) (*PMap, error) {
	pn := EmptyPMap(p.keyType, p.valueType)
	err := utils.CheckType(key, pn.keyType)
	if err != nil {
		return nil, err
	}
	err = utils.CheckType(value, pn.valueType)
	if err != nil {
		return nil, err
	}
	for k, v := range p.mapValue {
		pn.mapValue[k] = v
	}
	pn.mapValue[key] = value
	return pn, nil
}

func (p *PMap) Merge(q *PMap) (*PMap, error) {
	err := utils.CheckType(p.keyType, q.keyType)
	if err != nil {
		return nil, err
	}
	err = utils.CheckType(p.valueType, q.valueType)
	if err != nil {
		return nil, err
	}
	pn := EmptyPMap(p.keyType, p.valueType)
	for k, v := range p.mapValue {
		pn.mapValue[k] = v
	}
	for k, v := range q.mapValue {
		pn.mapValue[k] = v
	}
	return pn, nil
}

func (p *PMap) Get(key any) (any, error) {
	e := utils.CheckType(key, p.keyType)
	if e != nil {
		return nil, e
	}
	return p.mapValue[key], nil
}

// Delete removes the key-value pair identified by the key param.
// Returns a new PMap and the value that got deleted
func (p *PMap) Delete(key any) (*PMap, any, error) {
	pn := EmptyPMap(p.keyType, p.valueType)
	e := utils.CheckType(key, p.keyType)
	if e != nil {
		return nil, nil, e
	}
	var saved any
	for k, v := range p.mapValue {
		if k == key {
			saved = v
			continue
		}
		pn.mapValue[k] = v
	}
	return pn, saved, nil
}

// Keys retuns an array of the keys of PMap
func (p *PMap) Keys() []any {
	keys := make([]any, p.Length())
	i := 0
	for k := range p.mapValue {
		keys[i] = k
		i++
	}
	return keys
}

func (p *PMap) GetMap() map[any]any {
	pn := make(map[any]any)
	for k, v := range p.mapValue {
		pn[k] = v
	}
	return pn
}

func (p *PMap) Length() int {
	return len(p.mapValue)
}
