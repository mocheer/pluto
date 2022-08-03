package fn

func Panic(err error, msg string) {
	if err != nil {
		panic("[" + msg + "] " + err.Error())
	}
}
