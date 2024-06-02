package src

import (
	"fmt"
)

type Lexer struct {
	input          string
	character      byte
	start, current int //start = first char in lexeme being scanned--current = current character being scanned
}

/*
EOF         = "EOF"
*/

func ScanTokens(byteSource []byte) {
	/* var token Token */
	var tokenList []Token
	for index, value := range byteSource {
		switch value {
		case '{':
			tokenList = append(tokenList, createToken(RIGHT_BRACE, string(value)))
		case '}':
			tokenList = append(tokenList, createToken(LEFT_BRACE, string(value)))
		case '[':
			tokenList = append(tokenList, createToken(LEFT_PAREN, string(value)))
		case ']':
			tokenList = append(tokenList, createToken(RIGHT_PAREN, string(value)))
		case ',':
			tokenList = append(tokenList, createToken(COMMA, string(value)))
		case ':':
			tokenList = append(tokenList, createToken(COLON, string(value)))
		case '"':
			tokenList = append(tokenList, createToken(STRING, string(value)))
			/* stringToken(byteSource, index) */
			stringToken(byteSource[index:], index)
		}

	}
	tokenList = append(tokenList, Token{EOF, ""})
	fmt.Println(tokenList)
}

func createToken(tokenType string, Lexeme string) Token {
	return Token{TokenType: tokenType, Lexeme: string(Lexeme)}
}

func advance(byteSource []byte, index int) string {
	return string(byteSource[index+1])
}
func peek(byteSource []byte) string {
	if len(byteSource) <= 1 {
		return "" // rewrite tho throw throw an error here
	}
	return string(byteSource[1]) //byte slice is passed so to peek return index one value
	/* return string(byteSource[index+1]) */
}

func stringToken(byteSource []byte, startPosition int) {
	var stringBuildder string

	fmt.Println("len: ", len(byteSource), startPosition)
	for index := 0; peek(byteSource[startPosition:]) != `:`; index++ {
		/* for index := startPosition; peek(byteSource[index:]) != `:`; index++ { */
		if index >= len(byteSource[startPosition:]) {
			return
		}
		fmt.Println(index)
		stringBuildder += string(byteSource[index])
		fmt.Println(stringBuildder)
	}
	fmt.Println(stringBuildder)

}
