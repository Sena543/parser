package src

import (
	"fmt"
	"io"
)

type Parser struct {
	lexer        *Lexer
	currentToken Token
	nextToken    *Token
	parseErr     error
	//nextToken is my lookahead
}

func New(lexer *Lexer) *Parser {
	parser := &Parser{lexer: lexer}
	parser.getNextToken() //called twice to populate the two token fields in Parser struct
	parser.getNextToken() //cos on first call p.nextToken is nil, second call yeilds the scanned value
	return parser
}

func (p *Parser) getNextToken() error {
	if p.nextToken == nil {
		p.currentToken = Token{}
	} else {
		p.currentToken = *p.nextToken
	}
	token := p.lexer.ScanTokens()
	p.nextToken = &token

	if p.nextToken == nil {
		return fmt.Errorf("Error: nil next token")
	}

	return nil
}

func (p *Parser) ParserLoop(writer io.Writer) (string, error) {
	/* for !p.currentTokenIs(EOF) { */
	if p.currentToken.TokenType == LEFT_BRACE {
		if err := p.ParseObjects(); err != nil {
			return "invalid", err
		}

	} else if p.currentToken.TokenType == LEFT_PAREN {
		if err := p.ParseArray(); err != nil {
			return "invalid", err
		}
	} else {
		return "invalid", fmt.Errorf("json can only be array or object")
	}
	/* p.getNextToken() */
	/* } */

	if err := p.match(EOF); err != nil {
		return "invalid", fmt.Errorf("End of file expected")
	}
	return "valid", nil
}

func (p *Parser) ParseObjects() error {

	p.match(LEFT_BRACE)
	for !p.currentTokenIs(RIGHT_BRACE) {
		if err := p.match(STRING); err != nil {
			return fmt.Errorf("Expected string key but got %s", p.currentToken.Lexeme)
		}

		if err := p.match(COLON); err != nil {
			return err
		}

		if err := p.ParseValue(); err != nil {
			return err
		}

		if p.currentTokenIs(COMMA) && p.nextTokenIs(RIGHT_BRACE) {
			p.match(COMMA)
			p.match(RIGHT_BRACE)
			return fmt.Errorf("trailing comma")
		} else if p.currentTokenIs(COMMA) && p.nextTokenIs(STRING) {
			if err := p.match(COMMA); err != nil {
				return err
			}
		} else {
			break
		}
	}

	if err := p.match(RIGHT_BRACE); err != nil {
		return err
	}

	return nil
}

func (p *Parser) ParseValue() error {

	switch p.currentToken.TokenType {
	case STRING:
		return p.match(STRING)
	case LEFT_BRACE:
		return p.ParseObjects()
	case LEFT_PAREN:
		return p.ParseArray()
	case TRUE:
		return p.match(TRUE)
	case FALSE:
		return p.match(FALSE)
	case NUMBER:
		return p.match(NUMBER)
	case NULL:
		return p.match(NULL)
	}

	return fmt.Errorf("Value expected")
}

func (p *Parser) ParseArray() error {
	if err := p.match(LEFT_PAREN); err != nil {
		return err
	}
	/* fmt.Println(p.currentToken, p.nextToken) */
	for !p.currentTokenIs(RIGHT_PAREN) {

		if err := p.ParseValue(); err != nil {
			return err
		}

		if p.currentTokenIs(COMMA) && p.nextTokenIs(RIGHT_PAREN) {
			p.match(COMMA)
			p.match(RIGHT_PAREN)
			return fmt.Errorf("trailing comma")
		} else {

			if p.currentTokenIs(RIGHT_PAREN) {
				break
			}

			if err := p.match(COMMA); err != nil {
				return fmt.Errorf("Expected COMMA, found %v %v", p.currentToken.TokenType, p.currentToken.Lexeme)
			}
		}
	}

	if err := p.match(RIGHT_PAREN); err != nil {
		return err
	}
	return nil
}

func (p *Parser) match(expectedToken string) error {
	if p.currentToken.TokenType == expectedToken {
		err := p.getNextToken()
		if err != nil {
			return err
		}
		return nil
	}
	/* fmt.Println("expectedToken: ", expectedToken, "currentToken: ", p.currentToken, "nextToken: ", p.nextToken) */
	msg := fmt.Sprintf("Expected %s got %s", expectedToken, p.currentToken.TokenType)
	return fmt.Errorf(msg)
}

func (p *Parser) currentTokenIs(tk string) bool {
	return p.currentToken.TokenType == tk
}

func (p *Parser) nextTokenIs(tk string) bool {
	return p.nextToken.TokenType == tk
}
