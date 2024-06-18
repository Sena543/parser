package main

import (
	"fmt"
	"os"

	"parser/src"
)

func main() {
	readFile()
}

func readFile() {
	/* jsonPath := "./src/tests/step2/valid.json" */
	jsonPath := "./src/tests/step2/invalid.json"
	/* jsonPath := "./src/tests/step1/valid.json" */
	/* jsonPath := "./src/tests/step1/invalid.json" */
	fileBytes, error := os.ReadFile(jsonPath)
	if error != nil {
		fmt.Println("File reading error: ", error)
	}
	byteData := src.BeginScan(fileBytes)
	// byteData.ScanTokens()
	/* byteData.ScannerLoop() */
	parserInit := src.New(byteData)
	parserInit.ParserLoop()

}
