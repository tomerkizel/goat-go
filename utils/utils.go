package utils

import (
	"fmt"
	"reflect"
)

//CheckType returns an error if a and b of type any are not of the same real type. Returns nil otherwise
func CheckType(a, b any) error {
	if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
		return fmt.Errorf("vairable of type %v should be of type %v", reflect.TypeOf(a).Kind(), reflect.TypeOf(b).Kind())
	}
	return nil
}
