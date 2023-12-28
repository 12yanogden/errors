package errors

import (
	"fmt"
)

func Peek(err error) {
	if err != nil {
		Scream(err.Error())
	}
}

func Scream(str string) {
	panic(fmt.Sprintf("%s\n", str))
}
