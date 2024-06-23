package src

import (
	"fmt"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	readFile := func(filePath string) *Parser {
		t.Helper()
		fileBytes, error := os.ReadFile(filePath)
		if error != nil {
			fmt.Println("File reading error: ", error)
		}
		bytesData := BeginScan(fileBytes)
		return New(bytesData)

	}
	t.Run("step1: valid.json test", func(t *testing.T) {
		parse := readFile("./src/tests/step1/valid.json")

	})
	/* t.Run("step1: invalid.json test", func(t *testing.T) {
		 jsonPath := "./src/tests/step2/invalid.json"
	}) */
	/* t.Run("step1: invalid.json test", func(t *testing.T) {
		 jsonPath := "./src/tests/step2/invalid.json"
	}) */
	/* t.Run("step2: invalid.json test", func(t *testing.T) {
		 jsonPath := "./src/tests/step2/invalid.json"
	}) */
	/* t.Run("step2: invalid.json test", func(t *testing.T) {
		 jsonPath := "./src/tests/step2/invalid.json"
	}) */
}
