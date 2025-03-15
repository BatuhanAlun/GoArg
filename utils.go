package goarg

import "reflect"

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

func addFlag[T boolFlag | intFlag | stringFlag](flag T) {
	FlagList = append(FlagList, flag)
}
