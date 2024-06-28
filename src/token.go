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
	LEFT_PAREN      = "LEFT_PAREN"
	RIGHT_PAREN     = "RIGHT_PAREN"
	LEFT_BRACE      = "LEFT_BRACE"
	RIGHT_BRACE     = "RIGHT_BRACE"
	COMMA           = "COMMA"
	COLON           = "COLON"
	STRING          = "STRING"
	KEY             = "KEY"
	STRING_VALUE    = "STRING_VALUE"
	NUMBER          = "NUMBER"
	EOF             = "EOF"
	TRUE            = "TRUE"
	FALSE           = "FALSE"
	NULL            = "NULL"
	BACKSPACE       = "BACKSPACE"
	FORMFEED        = "FORMFEED"
	LINEFEED        = "LINEFEED"
	CARRIAGE_RETURN = "CARRIAGE_RETURN"
	HORIZONTAL_TAB  = "HORIZONTAL_TAB"
	SOLIDUS         = "SOLIDUS"
	REVERSE_SOLIDUS = "REVERSE_SOLIDUS"
)
