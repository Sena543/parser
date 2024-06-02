package src

type Token struct {
	TokenType string
	Lexeme    string
}

const (
	/* LEFT_PAREN  = "["
	RIGHT_PAREN = "]"
	LEFT_BRACE  = "}"
	RIGHT_BRACE = "{"
	COMMA       = ","
	EOF         = "EOF"
	*/
	LEFT_PAREN  = "LEFT_PAREN"
	RIGHT_PAREN = "RIGHT_PAREN"
	LEFT_BRACE  = "LEFT_BRACE"
	RIGHT_BRACE = "RIGHT_BRACE"
	COMMA       = "COMMA"
	COLON       = "COLON"
	STRING      = "STRING"
	EOF         = "EOF"
)
