package goarg

import (
	"fmt"
	"reflect"
)

// Gets all the variables and appends to FLAGLIST(Global Variable for tracking args)
func AddArg(value any, flagName string, defVal any, usageMessage string, strict bool) error {
	var pint *int
	var pstr *string
	var pbool *bool
	fType := findFlagType(value)
	if fType == reflect.TypeOf(pint) {
		retval := createIntFlag(value.(*int), flagName, defVal.(int), usageMessage, fType, strict)
		addFlag(&retval)
	} else if fType == reflect.TypeOf(pstr) {
		retval := createStringFlag(value.(*string), flagName, defVal.(string), usageMessage, fType, strict)
		addFlag(&retval)
	} else if fType == reflect.TypeOf(pbool) {
		retval := createBoolFlag(value.(*bool), flagName, defVal.(bool), usageMessage, fType, strict)
		addFlag(&retval)
	} else {
		return fmt.Errorf("unkown type")
	}
	return nil
}

func DeclareUsage(title, explanation string) {
	HelpMessage.Title = title
	HelpMessage.Explanation = explanation
	GetExamples()
	HelpMessage.Examples = UsageExampleSlice
}

// func AddUsageExample(exampleString string) {
// 	UsageExampleSlice.Examples = append(UsageExampleSlice.Examples, exampleString)
// }

func Parse() {
	checkHelp()
	argMap := createFlagMapValuePair() // map[flag]{value}
	mandatoryArgs := getMandatoryArgs()
	checkUnknownFlags(argMap)
	checkArgs(argMap, mandatoryArgs)
	giveValuesToArgs(argMap)
}
