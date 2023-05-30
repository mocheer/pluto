package fn

import gonanoid "github.com/matoous/go-nanoid/v2"

func Nanoid(l ...int) string {
	id, err := gonanoid.New(l...)
	if err != nil {
		panic(err)
	}
	return id
}
