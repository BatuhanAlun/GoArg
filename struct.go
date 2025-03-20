package goarg

import "reflect"

var FlagList []*any

type intFlag struct {
	FlagVar       *int
	FlagName      string
	FlagType      reflect.Type
	FlagDef       int
	FlagHelp      string
	FlagMandatory bool
}
type stringFlag struct {
	FlagVar       *string
	FlagName      string
	FlagType      reflect.Type
	FlagDef       string
	FlagHelp      string
	FlagMandatory bool
}
type boolFlag struct {
	FlagVar       *bool
	FlagName      string
	FlagType      reflect.Type
	FlagDef       bool
	FlagHelp      string
	FlagMandatory bool
}
