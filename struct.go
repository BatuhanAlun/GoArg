package goarg

import "reflect"

var FlagList []IFlag

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

type IFlag interface {
	GetHelp() string
	IsMandatory() bool
	GetFlagName() string
}

func (f *intFlag) GetHelp() string {
	return f.FlagHelp
}
func (f *intFlag) IsMandatory() bool {
	return f.FlagMandatory
}
func (f *intFlag) GetFlagName() string {
	return f.FlagName
}

func (f *stringFlag) GetHelp() string {
	return f.FlagHelp
}
func (f *stringFlag) IsMandatory() bool {
	return f.FlagMandatory
}
func (f *stringFlag) GetFlagName() string {
	return f.FlagName
}

func (f *boolFlag) GetHelp() string {
	return f.FlagHelp
}
func (f *boolFlag) IsMandatory() bool {
	return f.FlagMandatory
}
func (f *boolFlag) GetFlagName() string {
	return f.FlagName
}
