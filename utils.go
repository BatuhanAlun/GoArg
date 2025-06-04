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
	for _, flag := range FlagList {

		if valFromArgs, ok := argMap["-"+flag.GetFlagName()]; ok {

			if flag.GetFlagType().Elem().Kind() == reflect.Bool {

				if bVal, isBool := valFromArgs.(bool); isBool {
					flag.SetValue(bVal)
				} else if sVal, isString := valFromArgs.(string); isString {
					parsedBool, errr := strconv.ParseBool(sVal)
					if errr != nil {

						err(fmt.Errorf("invalid boolean value for flag %s: %v", flag.GetFlagName(), errr))
					}
					flag.SetValue(parsedBool)
				} else {

					err(fmt.Errorf("unexpected type for boolean flag %s value: %T", flag.GetFlagName(), valFromArgs))
				}
			} else {
				giveValueToPointers(flag, valFromArgs)
			}
		} else {

			flag.SetValue(flag.GetDefaultValue())
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

		if boolVal, ok := key.(bool); ok {
			value.SetValue(boolVal)
		} else if strVal, ok := key.(string); ok {

			parsedBool, errr := strconv.ParseBool(strVal)
			if errr != nil {
				err(fmt.Errorf("invalid boolean value for flag %s: %v", value.GetFlagName(), errr))
			}
			value.SetValue(parsedBool)
		} else {
			err(fmt.Errorf("unexpected type for boolean flag %s: %T", value.GetFlagName(), key))
		}
	default:
		err(fmt.Errorf("unknown type"))
	}
}

func checkHelp() {
	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("Title: %s\n", HelpMessage.Title)
		fmt.Printf("Explanation: %s\n", HelpMessage.Explanation)
		fmt.Println("Usages:")
		printExamples()
		os.Exit(1)
	}
}

func GetExamples() {
	for _, v := range FlagList {
		UsageExampleSlice.Examples = append(UsageExampleSlice.Examples, v.GetHelp())
	}
}

func printExamples() {
	for _, v := range HelpMessage.Examples.Examples {
		fmt.Println(v)
	}
}
