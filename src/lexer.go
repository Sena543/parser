package src

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
	// var tokenList []Token
	l.readChar()
	l.removeWhitespaces()
	/* fmt.Print(string(l.character)) */
	switch l.character {
	case '{':
		// tokenList = append(tokenList, createToken(RIGHT_BRACE, l.character))
		token = createToken(LEFT_BRACE, l.character)
	case '}':
		// tokenList = append(tokenList, createToken(LEFT_BRACE, l.character))
		token = createToken(RIGHT_BRACE, l.character)
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
	/* case 'f':
		token = Token{TokenType: FALSE, Lexeme: string(l.booleanToken())}
	case 't':
		token = Token{TokenType: TRUE, Lexeme: string(l.booleanToken())} */
	case '"': //check if is key or value in here
		// tokenList = append(tokenList, Token{TokenType: STRING, Lexeme: string(l.stringToken())})
		var tokenType string
		item := string(l.stringToken())
		if !l.atEnd() && l.input[l.current] == ':' {
			tokenType = KEY
		} else {
			tokenType = STRING_VALUE
			/* tokenType = VALUE */
		}
		token = Token{TokenType: tokenType, Lexeme: item}
	case '\000': //end of file
		token = createToken(EOF, '\000')
		//tok.Lexeme = ""
		//tok.Type = EOF
	default:
		if l.isDigit() {
			token = Token{TokenType: NUMBER, Lexeme: string(l.digitToken())}
		} else if l.isLetter() { //boolean check to extract value
			tokenValue := string(l.booleanToken())

			if tokenValue[0] == 't' {
				token = Token{TokenType: TRUE, Lexeme: tokenValue}
			} else {
				token = Token{TokenType: FALSE, Lexeme: tokenValue}
			}
		} else {
			token = Token{TokenType: NULL, Lexeme: "null"}
		}
	}

	/* fmt.Println(token) */
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
	return ('a' <= l.character && l.character <= 'z')
	/* return ('a' <= l.character && l.character <= 'z') || ('A' <= l.character && l.character <= 'Z') */
}

func (l *Lexer) booleanToken() []byte {

	l.start = l.current - 1 //firch char not being read so offet left by one
	for !l.atEnd() && l.isLetter() && l.peek() != ',' {
		l.readChar()
	}
	if l.atEnd() {
		panic("Comma separator or curly brace required after boolen value. found none")
	}

	l.readChar()
	return l.input[l.start : l.current-1]
	/* return l.input[l.start : l.current-1] */
}

func (l *Lexer) digitToken() []byte {

	l.start = l.current - 1 //firch char not being read so offet left by one
	for !l.atEnd() && l.isDigit() && l.peek() != ',' {
		l.readChar()
	}
	if l.atEnd() {
		panic("Comma separator or curly brace required after digits. found none")
	}

	//l.readChar() //read last quote
	return l.input[l.start : l.current-1]
}

func (l *Lexer) isDigit() bool {
	return l.character >= '0' && l.character <= '9'
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
