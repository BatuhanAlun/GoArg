package goarg

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
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

func giveValuesToArgs(argMap map[string]any) {
	var counter int
	if len(FlagList) != len(argMap) {
		err(fmt.Errorf("wrong arg number"))
	}
	for _, value := range FlagList {
		key, ok := argMap["-"+value.GetFlagName()]
		if !ok {
			counter++
		} else {
			giveValueToPointers(value, key)
		}
		if counter == len(FlagList) {
			err(fmt.Errorf("missing argument"))
		}
	}
}

func giveValueToPointers(value IFlag, key any) {
	v := value.GetFlagType().Elem().Kind()
	switch v {
	case reflect.Int:
		val, errr := strconv.Atoi(key.(string))
		if errr != nil {
			err(errr)
		}
		value.SetValue(val)
	case reflect.String:
		value.SetValue(key.(string))
	case reflect.Bool:
		value.SetValue(key)
	default:
		err(fmt.Errorf("unknown type"))
	}

}

func checkHelp() {
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println(HelpMessage)
		os.Exit(1)
	}
}
