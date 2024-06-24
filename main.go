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

func checkError(err error, msg string) {
	if err != nil {
		fmt.Println(fmt.Errorf(msg, err))
	}
}

func readFile() {
	stdIn, err := os.Stdin.Stat()
	checkError(err, "Error checking stdin")

	/* jsonPath := "./src/test_files/step2/valid.json" */
	/* jsonPath := "./src/test_files/step2/invalid.json" */
	/* jsonPath := "./src/tests_files/step3/valid.json" */
	/* jsonPath := "./src/tests_files/step3/ivalid.json" */
	/* jsonPath := "./src/tests_files/step3/valid.json" */
	/* jsonPath := "./src/tests_files/step3/invalid.json" */
	/* fileBytes, err := os.ReadFile(jsonPath) */
	/* checkError(err, "error reading file: ") */
	var fileDataPointer []byte
	if stdIn.Mode()&os.ModeCharDevice == 0 { //check if data is comming from stdin
		fileDataPointer, _ = io.ReadAll(os.Stdin)
	} else {
		filePath := os.Args[len(os.Args)-1]
		if _, err := os.Stat(filePath); err != nil { //check if path exists
			checkError(err, "File path error: ")
		}
		fileDataPointer, _ = os.ReadFile(filePath)
	}

	byteData := src.BeginScan(fileDataPointer)
	// byteData.ScanTokens()
	/* byteData.ScannerLoop() */
	parserInit := src.New(byteData)
	parserInit.ParserLoop(os.Stdout)
}
