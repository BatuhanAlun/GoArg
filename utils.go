package goarg

import (
	"os"
	"reflect"
)

func findFlagType(value interface{}) reflect.Type {
	flagType := reflect.TypeOf(value)
	return flagType
}

func createIntFlag(value *int, flagName string, defVal int, usageMessage string, fType reflect.Type) intFlag {
	return intFlag{FlagVar: value, FlagName: flagName, FlagDef: defVal, FlagHelp: usageMessage, FlagType: fType}
}

func createStringFlag(value *string, flagName string, defVal string, usageMessage string, fType reflect.Type) stringFlag {
	return stringFlag{FlagVar: value, FlagName: flagName, FlagDef: defVal, FlagHelp: usageMessage, FlagType: fType}
}

func createBoolFlag(value *bool, flagName string, defVal bool, usageMessage string, fType reflect.Type) boolFlag {
	return boolFlag{FlagVar: value, FlagName: flagName, FlagDef: defVal, FlagHelp: usageMessage, FlagType: fType}
}

func addFlag(flag any) {
	FlagList = append(FlagList, &flag)
}

func createRFlagNameList() []string {
	var rFlagNameList []string
	for _, v := range os.Args[1:] {
		rFlagName := v[1:]
		rFlagNameList = append(rFlagNameList, rFlagName)
	}
	return rFlagNameList
}
