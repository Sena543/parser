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

	t.Run("step1:", func(t *testing.T) {
		step1 := []struct {
			path string
			want string
		}{
			{path: "./tests_files/step1/valid.json", want: "valid\n"},
			{path: "./tests_files/step1/invalid.json", want: "Error: file empty"},
			/* {path: "./tests_files/step1/invalid.json", want: "Error: file empty"}, */
		}
		for _, value := range step1 {
			t.Run(value.path, func(t *testing.T) {

				buffer := bytes.Buffer{}

				parse := readFile(t, value.path)
				parse.ParserLoop(&buffer)
				got := buffer.String()
				want := value.want

				if got != want {
					t.Errorf("got %q want %q \n", got, want)
				}
				buffer.Reset()

			})

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
