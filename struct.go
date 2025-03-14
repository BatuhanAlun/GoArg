package goarg

type intFlag struct {
	FlagVar  *int
	FlagName string
	FlagType string
	FlagDef  string
	FlagHelp string
}
type stringFlag struct {
	FlagVar  *string
	FlagName string
	FlagType string
	FlagDef  string
	FlagHelp string
}
type boolFlag struct {
	FlagVar  *bool
	FlagName string
	FlagType string
	FlagDef  string
	FlagHelp string
}
