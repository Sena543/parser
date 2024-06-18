package src

import "fmt"

type Parser struct {
	lexer                   *Lexer
	currentToken, nextToken Token
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

func (p *Parser) ParserLoop() {
	for p.currentToken.TokenType != EOF {
		if p.currentToken.TokenType == LEFT_BRACE {
			p.ParseObjects()
		}
		fmt.Println(p.currentToken, p.nextToken)
		/* p.getNextToken() */
	}

	if p.currentToken.TokenType == EOF {
		fmt.Println("Input file is valid")
	}
}

func (p *Parser) ParseString() {

}

func (p *Parser) ParseObjects() {

	if p.nextTokenIs(RIGHT_BRACE) {
		p.match(LEFT_BRACE)
		p.match(RIGHT_BRACE)
		return
	}

	if p.nextTokenIs(KEY) {
		p.match(LEFT_BRACE)
		p.match(KEY)
		p.match(COLON)
		p.ParseValue()
		return
	}
}

func (p *Parser) ParseValue() {
	p.getNextToken()
	p.getNextToken()
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
	if p.currentToken.TokenType == expectedToken {
		p.getNextToken()
		return
	}
	msg := fmt.Sprintf("Expected %s got %s", p.currentToken.TokenType, expectedToken)
	panic(msg)
}

func (p *Parser) currentTokenIs(tk string) bool {
	return p.currentToken.TokenType == tk
}

func (p *Parser) nextTokenIs(tk string) bool {
	return p.nextToken.TokenType == tk
}
