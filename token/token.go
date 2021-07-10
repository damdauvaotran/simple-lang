package token

type Type string

type Token struct {
	Type    Type   // Type type of token
	Literal string // Literal value of token
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	DIVIDE   = "/"
	MULTIPLE = "*"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]Type{
	"fn":  FUNCTION,
	"let": LET,
}

// LookUpIdent get type of the identifier is (keyword or variable name)
func LookUpIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// New create new token
func New(tokenType Type, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

