package src

import "fmt"

type Parser struct {
	lexer                   *Lexer
	currentToken, nextToken Token
}

func New(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.getToken() //called twice to populate the two token fields in Parser struct
	parser.getToken() //cos on first call p.nextToken is nil, second call yeilds the scanned value
	return parser
}

func (p *Parser) getToken() {
	p.currentToken = p.nextToken
	p.nextToken = p.lexer.ScanTokens()
}

func (p *Parser) ParserLoop() {
	for p.currentToken.TokenType != EOF {
		if p.currentToken.TokenType == LEFT_BRACE {
			p.ParseObjects()
		}
		fmt.Println(p.currentToken, p.nextToken)
		p.getToken()
	}

}

func (p *Parser) ParseString() {

}

func (p *Parser) ParseObjects() {
	if p.expect(RIGHT_BRACE) {
		p.getToken()
		return
	}

	if p.expect(KEY) {
		p.getToken()
		/* return */
	}

	if p.expect(COLON) {
		p.getToken()
	}
}

func (p *Parser) ParseValue() {
}

// checks the what we expect the nextToken to be
// tk string type same as Token.TokenType
func (p *Parser) expect(tk string) bool {
	if !p.nextTokenIs(tk) {
		return false
	}
	return true
}

func (p *Parser) currentTokenIs(tk string) bool {
	return p.currentToken.TokenType == tk
}

func (p *Parser) nextTokenIs(tk string) bool {
	return p.nextToken.TokenType == tk
}
