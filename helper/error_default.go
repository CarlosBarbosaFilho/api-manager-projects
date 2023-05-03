package helper

func ErrorDefault(err error) {
	if err != nil {
		panic(err)
	}
}
