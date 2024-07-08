package src

import (
	"fmt"
	"io"
	"os"
)

type Parser struct {
	lexer                   *Lexer
	currentToken, nextToken Token
	stdout                  io.Writer
	//nextToken is my lookahead
}

func New(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.stdout = os.Stdout //print to stdout by default
	parser.getNextToken()     //called twice to populate the two token fields in Parser struct
	parser.getNextToken()     //cos on first call p.nextToken is nil, second call yeilds the scanned value
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
		} else if p.currentToken.TokenType == LEFT_PAREN {
			p.ParseArray()
		} else {
			p.ParseValue()
		}
		p.getNextToken()
	}

	if p.currentTokenIs(EOF) {
		fmt.Fprintln(writer, "valid")
	}
}

func (p *Parser) ParseObjects() {

	p.match(LEFT_BRACE)
	for !p.currentTokenIs(RIGHT_BRACE) {
		p.match(KEY)
		p.match(COLON)
		if p.nextTokenIs(COMMA) {
			p.ParseValue()
			if p.nextTokenIs(RIGHT_BRACE) {
				p.match(COMMA)
				/* panic("trailing comma") */
				p.parserError("trailing comma")
			} else {
				p.match(COMMA)
			}
		} else {
			p.ParseValue()
		}
	}

	p.match(RIGHT_BRACE)
	if p.nextTokenIs(KEY) {
		p.match(COMMA)
	}
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
	default:
		p.parserError("default: illegal value")
		p.IllegalValue(p.currentToken)
	}

}

func (p *Parser) IllegalValue(expectedToken Token) {

	msg := fmt.Sprintf("Unexpected value %s", expectedToken.Lexeme)
	fmt.Println(p.stdout, msg)
	panic(msg)
}

func (p *Parser) parserError(errMsg string) {
	fmt.Println(p.stdout, errMsg)
	os.Exit(1)
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
	/* fmt.Println(expectedToken, p.currentToken.Lexeme, "next token:", p.nextToken.Lexeme, p.nextToken.TokenType) */
	if p.currentToken.TokenType == expectedToken {
		p.getNextToken()
		return
	}
	msg := fmt.Sprintf("Expected %s got %s", expectedToken, p.currentToken.TokenType)
	p.parserError(msg)
	/* panic(msg) */
}

func (p *Parser) currentTokenIs(tk string) bool {
	return p.currentToken.TokenType == tk
}

func (p *Parser) nextTokenIs(tk string) bool {
	return p.nextToken.TokenType == tk
}
