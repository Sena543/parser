package src

import (
	"fmt"
	"io"
)

type Parser struct {
	lexer                   *Lexer
	currentToken, nextToken Token
	//nextToken is my lookahead
}

func New(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.getNextToken() //called twice to populate the two token fields in Parser struct
	parser.getNextToken() //cos on first call p.nextToken is nil, second call yeilds the scanned value
	return parser
}

func (p *Parser) getNextToken() {
	p.currentToken = p.nextToken
	p.nextToken = p.lexer.ScanTokens()
}

func (p *Parser) ParserLoop(writer io.Writer) {
	for p.currentToken.TokenType != EOF {
		if p.currentToken.TokenType == LEFT_BRACE {
			p.ParseObjects()
		}
		p.getNextToken()
	}

	if p.currentTokenIs(EOF) {
		fmt.Fprintln(writer, "Input file is valid")
		/* fmt.Println("Input file is valid") */
	}
}

func (p *Parser) ParseObjects() {

	if p.nextTokenIs(RIGHT_BRACE) {
		p.match(LEFT_BRACE)
		p.match(RIGHT_BRACE)
		return
	}

	p.match(LEFT_BRACE)
	for !p.currentTokenIs(RIGHT_BRACE) {
		p.match(KEY)
		p.match(COLON)
		p.ParseValue()

		if p.nextTokenIs(KEY) {
			/* if p.currentTokenIs(COMMA) { */
			p.match(COMMA)
		}

		/* 	else if p.nextTokenIs(RIGHT_BRACE) && p.currentTokenIs(COMMA) {
			p.match(COMMA) //match to make sure there is a trailing comma
			panic("Trailing comma")
		} */

	}

	p.match(RIGHT_BRACE)
}

func (p *Parser) ParseValue() {

	switch p.currentToken.TokenType {
	case STRING_VALUE:
		p.match(STRING_VALUE)
	case LEFT_BRACE:
		p.ParseObjects()
	case LEFT_PAREN:
		p.ParseArray()
	case TRUE:
		p.match(TRUE)
	case FALSE:
		p.match(FALSE)
	case NUMBER:
		p.match(NUMBER)
	case NULL:
		p.match(NULL)
	}

}

func (p *Parser) ParseArray() {

	p.match(LEFT_PAREN)
	for !p.currentTokenIs(RIGHT_PAREN) {
		p.ParseValue()
		if p.nextTokenIs(COMMA) {
			p.match(COMMA)
		}
	}
	p.match(RIGHT_PAREN)
}

// checks the what we expect the nextToken to be
// tk string type same as Token.TokenType
func (p *Parser) expect(tk string) bool {
	if !p.nextTokenIs(tk) {
		return false
	}
	return true
}

func (p *Parser) match(expectedToken string) {
	fmt.Println(expectedToken, p.currentToken.Lexeme)
	if p.currentToken.TokenType == expectedToken {
		/* 	if p.nextTokenIs(ILLEGAL) {
			panic("illegal value")
		} */
		p.getNextToken()
		return
	}
	msg := fmt.Sprintf("Expected %s got %s", expectedToken, p.currentToken.TokenType)
	panic(msg)
}

func (p *Parser) currentTokenIs(tk string) bool {
	return p.currentToken.TokenType == tk
}

func (p *Parser) nextTokenIs(tk string) bool {
	return p.nextToken.TokenType == tk
}
