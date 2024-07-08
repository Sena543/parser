package src

import (
	"fmt"
	"io"
)

func CheckError(writer io.Writer, err error, msg string) {
	if err != nil {
		fmt.Println(writer, fmt.Errorf(msg, err))
	}
}
