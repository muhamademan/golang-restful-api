package helper

// function ini digunakan ketika terdapat panic err yang berulang / double panic
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
