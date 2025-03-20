package goarg

import (
	"fmt"
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

func addFlag(flag IFlag) {
	FlagList = append(FlagList, flag)
}

func createFlagMapValuePair() map[string]any {
	flagAndValues := make(map[string]any)

	for index, v := range os.Args[1:] {
		if v[0] == '-' && index+1 == len(os.Args[1:]) {
			flagAndValues[v] = true
		} else if v[0] == '-' && os.Args[1:][index+1][0] == '-' {
			flagAndValues[v] = true
		} else if v[0] == '-' && os.Args[1:][index+1][0] != '-' {
			flagAndValues[v] = os.Args[1:][index+1]
		}
	}
	return flagAndValues
}

func getMandatoryArgs() []IFlag {
	var mandatoryArgs []IFlag
	for _, v := range FlagList {
		if v.IsMandatory() {
			mandatoryArgs = append(mandatoryArgs, v)
		}
	}
	return mandatoryArgs
}

func checkArgs(argMap map[string]any, mandatoryArgs []IFlag) {
	if len(argMap) < len(mandatoryArgs) {
		err(fmt.Errorf("please be sure to insert Mandatory Arguments"))
	}
	for _, v := range mandatoryArgs {
		_, ok := argMap["-"+v.GetFlagName()]
		if !ok {
			err(fmt.Errorf("missing mandatory flag"))
		}
	}
}
