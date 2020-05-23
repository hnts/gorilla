package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifer + literal
	IDENT = "IDENT"
	INT   = "INT"

	// oparator
	ASSIGN = "="
	PLUS   = "+"

	// delimiter
	COMMA     = ","
	SEMICOLON = ";" 

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)