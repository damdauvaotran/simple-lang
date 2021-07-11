package token

type Type string

type Token struct {
	Type    Type   // Type type of token
	Literal string // Literal value of token
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	MODULO   = "%"

	// Comparator
	EQ     = "=="
	NOT_EQ = "!="
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = ">="
	BANG   = "!"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	FOR      = "FOR"
	WHILE    = "WHILE"
	NEW      = "NEW"
	RETURN   = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"for":    FOR,
	"while":  WHILE,
	"new":    NEW,
	"return": RETURN,
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
