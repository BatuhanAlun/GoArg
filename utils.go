package goarg

import (
	"os"
	"reflect"
)

func findFlagType(value interface{}) reflect.Type {
	flagType := reflect.TypeOf(value)
	return flagType
}

func createIntFlag(value *int, flagName string, defVal int, usageMessage string, fType reflect.Type, strict bool) intFlag {
	return intFlag{FlagVar: value, FlagName: flagName, FlagDef: defVal, FlagHelp: usageMessage, FlagType: fType, FlagMandatory: strict}
}

func createStringFlag(value *string, flagName string, defVal string, usageMessage string, fType reflect.Type, strict bool) stringFlag {
	return stringFlag{FlagVar: value, FlagName: flagName, FlagDef: defVal, FlagHelp: usageMessage, FlagType: fType, FlagMandatory: strict}
}

func createBoolFlag(value *bool, flagName string, defVal bool, usageMessage string, fType reflect.Type, strict bool) boolFlag {
	return boolFlag{FlagVar: value, FlagName: flagName, FlagDef: defVal, FlagHelp: usageMessage, FlagType: fType, FlagMandatory: strict}
}

func addFlag(flag any) {
	FlagList = append(FlagList, &flag)
}

func createFlagMapValuePair() map[string]any {
	flagAndValues := make(map[string]any)

	for index, v := range os.Args[1:] {
		if v[1] == '-' && os.Args[1:][index+1][1] != '-' {
			flagAndValues[v] = os.Args[1:][index+1]
		} else if v[1] == '-' && os.Args[1:][index+1][1] == '-' {
			flagAndValues[v] = true
		}
	}
	return flagAndValues
}
