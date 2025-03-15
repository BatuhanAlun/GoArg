package goarg

func err(err error) {
	if err != nil {
		panic(err)
	}
}
