package src

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	readFile := func(t testing.TB, filePath string) *Parser {
		t.Helper()
		fileBytes, error := os.ReadFile(filePath)
		if error != nil {
			fmt.Println("File reading error: ", error)
		}
		bytesData := BeginScan(fileBytes)
		return New(bytesData)

	}
	t.Run("step1: valid.json test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		parse := readFile(t, "./tests_files/step1/valid.json")
		parse.ParserLoop(&buffer)
		got := buffer.String()
		want := "Input file is valid\n"

		if got != want {
			t.Errorf("got %q want %q \n", got, want)
		}

	})
	/* t.Run("step1: invalid.json test", func(t *testing.T) {
		 jsonPath := "./src/test_files/step2/invalid.json"
	}) */
	/* t.Run("step1: invalid.json test", func(t *testing.T) {
		 jsonPath := "./src/test_files/step2/invalid.json"
	}) */
	/* t.Run("step2: invalid.json test", func(t *testing.T) {
		 jsonPath := "./src/test_files/step2/invalid.json"
	}) */
	/* t.Run("step2: invalid.json test", func(t *testing.T) {
		 jsonPath := "./src/test_files/step2/invalid.json"
	}) */
}
