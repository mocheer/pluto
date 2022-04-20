package mocker

func NameLast() string {
	return NameLastArray[Intn(len(NameLastArray))]
}

func NameFirst() string {
	return NameFirstArray[Intn(len(NameFirstArray))] + NameFirstArray[Intn(len(NameFirstArray))]
}

func Name() string {
	return NameLast() + NameFirst()
}
