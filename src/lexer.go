package src

import "fmt"

type Lexer struct {
	// input          string
	input          []byte
	character      byte
	start, current int //start = first char in lexeme being scanned--current = current character being scanned
}

func BeginScan(inputSource []byte) *Lexer {
	lexer := Lexer{input: inputSource}
	return &lexer
}

func (l *Lexer) ScannerLoop() {
	for l.current < len(l.input) {
		fmt.Println("1-: ", l.current, len(l.input))
		l.ScanTokens()
		// l.current++
		fmt.Println("2-: ", l.current, len(l.input))
	}
}

func (l *Lexer) ScanTokens() {
	var token Token
	// var tokenList []Token
	l.readChar()
	l.removeWhitespaces()
	switch l.character {
	case '{':
		// tokenList = append(tokenList, createToken(RIGHT_BRACE, l.character))
		token = createToken(RIGHT_BRACE, l.character)
	case '}':
		// tokenList = append(tokenList, createToken(LEFT_BRACE, l.character))
		token = createToken(LEFT_BRACE, l.character)
		fmt.Println("this line ran")
	case '[':
		// tokenList = append(tokenList, createToken(LEFT_PAREN, l.character))
		token = createToken(LEFT_PAREN, l.character)
	case ']':
		// tokenList = append(tokenList, createToken(RIGHT_PAREN, l.character))
		token = createToken(RIGHT_PAREN, l.character)
	case ',':
		// tokenList = append(tokenList, createToken(COMMA, l.character))
		token = createToken(COMMA, l.character)
	case ':':
		// tokenList = append(tokenList, createToken(COLON, l.character))
		token = createToken(COLON, l.character)
	case '"': //check if is key or value in here
		// tokenList = append(tokenList, Token{TokenType: STRING, Lexeme: string(l.stringToken())})
		token = Token{TokenType: STRING, Lexeme: string(l.stringToken())}
	}

	fmt.Println(token)
	// fmt.Println(tokenList)
}

func createToken(tokenType string, literal byte) Token {
	return Token{TokenType: tokenType, Lexeme: string(literal)}
}

func (l *Lexer) readChar() {

	l.current++
	if l.current > len(l.input) {
		// fmt.Println("called readchar", l.atEnd())
		return
	}
	l.character = l.input[l.current]

	// l.character = l.input[l.current-1]
}

func (l *Lexer) peek() byte {
	if l.atEnd() {
		return 0
	}
	return l.input[l.current]
}

func (l *Lexer) atEnd() bool {
	// fmt.Println("something: ", len(l.input), l.current)
	return l.current >= len(l.input)
}

func (l *Lexer) removeWhitespaces() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
}

// func (l *Lexer) isLetter() bool {
// 	return ('a' <= l.character && l.character <= 'z') || ('A' <= l.character && l.character <= 'Z')
// }

func (l *Lexer) stringToken() []byte {
	// fmt.Println(l.start, l.current)
	l.start = l.current
	for l.peek() != '"' && !l.atEnd() {
		l.readChar()
		fmt.Println("string tokens")
	}

	if l.atEnd() {
		panic("Unterminated string")
	}

	fmt.Println("first: ", string(l.input[l.start:l.current]))
	l.readChar() //last quote
	fmt.Println("second: ", string(l.input[l.start:l.current]))
	return l.input[l.start:l.current]
}
