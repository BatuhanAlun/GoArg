package goarg

func AddArg(value any, flagName, defVal, usageMessage string) {
	createFlag(value, defVal, usageMessage)
}
