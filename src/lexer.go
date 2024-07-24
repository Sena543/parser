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
		l.ScanTokens()
	}
}

func (l *Lexer) ScanTokens() Token {

	var token Token

	l.readChar()
	l.removeWhitespaces()
	switch l.character {
	case '{':
		token = createToken(LEFT_BRACE, l.character)
	case '}':
		token = createToken(RIGHT_BRACE, l.character)
	case '[':
		token = createToken(LEFT_PAREN, l.character)
	case ']':
		token = createToken(RIGHT_PAREN, l.character)
	case ',':
		token = createToken(COMMA, l.character)
	case ':':
		token = createToken(COLON, l.character)
	case '"': //check if is key or value in here
		var tokenType string
		item := string(l.stringToken())
		if !l.atEnd() && l.input[l.current] == ':' {
			tokenType = KEY
		} else {
			tokenType = STRING_VALUE
		}
		token = Token{TokenType: tokenType, Lexeme: item}
	case '\000': //end of file
		token = createToken(EOF, '\000')
	default:
		if l.isDigit() {
			token = Token{TokenType: NUMBER, Lexeme: string(l.digitToken())}
		} else if l.isLetter() { //boolean check to extract true or false value
			tokenValue := string(l.booleanToken())

			if tokenValue == "true" {
				/* if tokenValue[0] == 't' { */
				token = Token{TokenType: TRUE, Lexeme: tokenValue}
				/* } else if tokenValue[0] == 'f' { */
			} else if tokenValue == "false" {
				token = Token{TokenType: FALSE, Lexeme: tokenValue}
			} else if tokenValue == "null" { //used string check instead of char cos value might start with and n
				//must do same for above checks
				token = Token{TokenType: NULL, Lexeme: "null"}
			} else {

				/* fmt.Println("illegal token :", tokenValue) */
				token = Token{TokenType: ILLEGAL, Lexeme: tokenValue}
				/* token = Token{TokenType: NULL, Lexeme: "null"} */
			}
		} /* else {
			token = Token{TokenType: NULL, Lexeme: "null"}
		} */
	}

	/* fmt.Println("lexer: ", token) */
	return token
}

func createToken(tokenType string, literal byte) Token {
	return Token{TokenType: tokenType, Lexeme: string(literal)}
}

func (l *Lexer) readChar() {

	if l.atEnd() { //handles end of input
		/* l.character = 0 */
		l.character = '\000'
		return
	}
	l.character = l.input[l.current]
	l.current++
}

func (l *Lexer) peek() byte {
	if l.atEnd() {
		return 0
	}
	return l.input[l.current]
}

func (l *Lexer) atEnd() bool {
	return l.current >= len(l.input)
}

func (l *Lexer) removeWhitespaces() {
	for l.character == ' ' || l.character == '\t' || l.character == '\n' || l.character == '\r' {
		l.readChar()
	}
}

func (l *Lexer) isLetter() bool {
	return ('a' <= l.character && l.character <= 'z') || ('A' <= l.character && l.character <= 'Z')
}

func (l *Lexer) booleanToken() []byte {

	l.start = l.current - 1                             //first char not being read so offet left by one
	for !l.atEnd() && l.isLetter() && l.peek() != ',' { //potential bug here what happens if l.peek()==}/{ etc
		l.readChar()
	}
	return l.input[l.start:l.current]
	/* return l.input[l.start : l.current-1] */
}

func (l *Lexer) digitToken() []byte {

	if l.input[l.current-2] == '-' {
		l.start = l.current - 2 //first char not being read so offet left by two to include - character
	} else {
		l.start = l.current - 1 //first char not being read so offet left by one
	}

	for !l.atEnd() && l.isDigit() && l.peek() != ',' {
		l.readChar()
	}

	if l.character == '.' { //read the decimal point
		l.readChar()
	}
	for !l.atEnd() && l.isDigit() && l.peek() != ',' { // read rest of  numbers
		l.readChar()
	}

	if l.character == 'E' || l.character == 'e' {
		l.readChar()
		if l.character == '+' || l.character == '-' {
			l.readChar()
		}
		for !l.atEnd() && l.isDigit() && l.peek() != ',' { // read rest of  numbers
			l.readChar()
		}
	}

	return l.input[l.start:l.current]
}

func (l *Lexer) isDigit() bool {
	return (l.character >= '0' && l.character <= '9')
	/* return (l.character >= '0' && l.character <= '9') || (l.character == '-' && l.peek() >= '0' && l.peek() <= '9') */
}

func (l *Lexer) stringToken() []byte {

	l.start = l.current
	for !l.atEnd() && l.input[l.current] != '"' {
		/* for !l.atEnd() && l.peek() != '"' { */
		l.readChar()
	}
	if l.atEnd() {
		panic("Unterminated string")
	}

	l.readChar() //read last quote
	return l.input[l.start : l.current-1]
}
