package utils

import (
	"fmt"
	"reflect"
)

func checkType(a, b any) error {
	if reflect.TypeOf(a).Kind() != reflect.TypeOf(b).Kind() {
		return fmt.Errorf("key of type %v should be of type %v", reflect.TypeOf(a).Kind(), reflect.TypeOf(b).Kind())
	}
	return nil
}
