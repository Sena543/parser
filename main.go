package main

import (
	"fmt"
	"io"
	"os"

	"parser/src"
)

func main() {
	readFile()
}

/* func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(fmt.Errorf(msg, err))
	}
} */

func readFile() {
	stdIn, err := os.Stdin.Stat()
	src.CheckError(os.Stdout, err, "Error checking stdin")

	var fileDataPointer []byte
	if stdIn.Mode()&os.ModeCharDevice == 0 { //check if data is comming from stdin
		fileDataPointer, _ = io.ReadAll(os.Stdin)
	} else {
		filePath := os.Args[len(os.Args)-1]
		if _, err := os.Stat(filePath); err != nil { //check if path exists
			src.CheckError(os.Stdout, err, "File path error: ")
		}
		fileDataPointer, _ = os.ReadFile(filePath)
	}

	if len(fileDataPointer) < 1 {
		fmt.Println("Error: file empty")
		return
	}

	byteData := src.BeginScan(fileDataPointer)
	byteData.ScannerLoop()
	parserInit := src.New(byteData)
	res, err := parserInit.ParserLoop(os.Stdout)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}
