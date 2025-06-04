package goarg

import (
	"fmt"
	"reflect"
)

var FlagList []IFlag
var UsageExampleSlice UsageExamples
var HelpMessage Usage

type UsageExamples struct {
	Examples []string
}
type Usage struct {
	Title       string
	Explanation string
	Examples    UsageExamples
}

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
	GetFlagType() reflect.Type
	SetValue(newVal any)
	GetDefaultValue() any
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
func (f *intFlag) GetFlagPointer() any {
	return f.FlagVar
}
func (f *intFlag) GetFlagType() reflect.Type {
	return f.FlagType
}
func (f *intFlag) SetValue(newVal any) {
	if val, ok := newVal.(int); ok {
		*f.FlagVar = val
	} else {
		err(fmt.Errorf("setting wrong type"))
	}
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
func (f *stringFlag) GetFlagPointer() any {
	return f.FlagVar
}
func (f *stringFlag) GetFlagType() reflect.Type {
	return f.FlagType
}
func (f *stringFlag) SetValue(newVal any) {
	if val, ok := newVal.(string); ok {
		*f.FlagVar = val
	}
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
func (f *boolFlag) GetFlagPointer() any {
	return f.FlagVar
}
func (f *boolFlag) GetFlagType() reflect.Type {
	return f.FlagType
}
func (f *boolFlag) SetValue(newVal any) {
	if val, ok := newVal.(bool); ok {
		*f.FlagVar = val
	}
}

func (f *intFlag) GetDefaultValue() any {
	return f.FlagDef
}

func (f *stringFlag) GetDefaultValue() any {
	return f.FlagDef
}

func (f *boolFlag) GetDefaultValue() any {
	return f.FlagDef
}
