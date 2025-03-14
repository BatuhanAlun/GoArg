package goarg

import "reflect"

func findFlagType(value interface{}) reflect.Type {
	flagType := reflect.TypeOf(value)
	return flagType
}
