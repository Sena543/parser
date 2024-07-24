package src

import (
	"bytes"
	"errors"
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
			{path: "./tests_files/step1/valid.json", want: "valid"},
			/* {path: "./tests_files/step1/invalid.json", want: "Error: file empty"}, */
		}
		for _, value := range step1 {
			t.Run(value.path, func(t *testing.T) {

				buffer := bytes.Buffer{}

				parse := readFile(t, value.path)
				got, _ := parse.ParserLoop(&buffer)
				/* got := buffer.String() */
				want := value.want

				if got != want {
					t.Errorf("got %q want %q \n", got, want)
				}
				buffer.Reset()
			})
		}
	})
	t.Run("step2: ", func(t *testing.T) {
		step2 := []struct {
			path    string
			want    string
			message error
		}{
			{path: "./tests_files/step2/valid.json", want: "valid", message: nil},
			{path: "./tests_files/step2/invalid.json", want: "invalid", message: errors.New("trailing comma")},
			{path: "./tests_files/step2/valid2.json", want: "valid", message: nil},
			{path: "./tests_files/step2/invalid2.json", want: "invalid", message: errors.New("Expected KEY got ILLEGAL")},
		}
		for _, value := range step2 {
			t.Run(value.path, func(t *testing.T) {

				buffer := bytes.Buffer{}

				parse := readFile(t, value.path)
				/* got, errMsg := parse.ParserLoop(&buffer) */
				got, _ := parse.ParserLoop(&buffer)
				want := value.want

				if got != want {
					t.Errorf("got %q want %q \n", got, want)
				}
				/* 		if value.message != errMsg {
					t.Errorf("got %q want %q \n", got, want)
				}
				*/buffer.Reset()
			})
		}
	})
	t.Run("step3: ", func(t *testing.T) {
		step3 := []struct {
			path    string
			want    string
			message error
		}{
			{path: "./tests_files/step3/valid.json", want: "valid", message: nil},
			{path: "./tests_files/step3/invalid.json", want: "invalid", message: errors.New("Value expected")},
		}
		for _, value := range step3 {
			t.Run(value.path, func(t *testing.T) {

				buffer := bytes.Buffer{}
				parse := readFile(t, value.path)
				got, el := parse.ParserLoop(&buffer)
				want := value.want

				fmt.Println("got", el)
				if got != want {
					t.Errorf("got %q want %q \n", got, want)
				}
			})
		}

	})
	t.Run("step4: ", func(t *testing.T) {
		step4 := []struct {
			path    string
			want    string
			message error
		}{
			{path: "./tests_files/step4/valid.json", want: "valid", message: nil},
			{path: "./tests_files/step4/valid2.json", want: "valid", message: nil},
			{path: "./tests_files/step4/invalid.json", want: "invalid", message: errors.New("trailing comma")},
		}
		for _, value := range step4 {
			t.Run(value.path, func(t *testing.T) {

				buffer := bytes.Buffer{}
				parse := readFile(t, value.path)
				got, _ := parse.ParserLoop(&buffer)
				want := value.want

				if got != want {
					t.Errorf("got %q want %q \n", got, want)
				}
			})
		}

	})
}
