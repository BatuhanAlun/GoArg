package goarg

import (
	"fmt"
	"os"
	"reflect"
)

// Gets all the variables and appends to FLAGLIST(Global Variable for tracking args)
func AddArg(value any, flagName string, defVal any, usageMessage string) error {
	var pint *int
	var pstr *string
	var pbool *bool
	fType := findFlagType(value)
	if fType == reflect.TypeOf(pint) {
		retval := createIntFlag(value.(*int), flagName, defVal.(int), usageMessage, fType)
		addFlag(retval)
	} else if fType == reflect.TypeOf(pstr) {
		retval := createStringFlag(value.(*string), flagName, defVal.(string), usageMessage, fType)
		addFlag(retval)
	} else if fType == reflect.TypeOf(pbool) {
		retval := createBoolFlag(value.(*bool), flagName, defVal.(bool), usageMessage, fType)
		addFlag(retval)
	} else {
		return fmt.Errorf("unkown type")
	}
	return nil
}

func Parse() {
	var args []string = os.Args[1:]
	rFlagNameList := createRFlagNameList()
	for _, v := range args {
		rFlagName := v[1:]
	}

}
