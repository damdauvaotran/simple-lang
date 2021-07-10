package lexer

import (
	"amonkey/token"
	"fmt"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar read the next character in input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// NextToken read the next token in input, skip the whitespace
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()

	switch l.ch {
	case '=':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.ASSIGN, l.ch)
	case ';':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.SEMICOLON, l.ch)
	case '(':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.LPAREN, l.ch)
	case ')':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.RPAREN, l.ch)
	case ',':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.COMMA, l.ch)
	case '+':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.PLUS, l.ch)
	case '-':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.MINUS, l.ch)
	case '*':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.MULTIPLE, l.ch)
	case '/':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.DIVIDE, l.ch)
	case '}':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.RBRACE, l.ch)
	case '{':
		fmt.Printf("scanned %v \n", string(l.ch))
		tok = token.New(token.LBRACE, l.ch)
	case 0:
		fmt.Printf("scanned %v \n", "EOF")
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			var ident = l.readIdentifier()
			tok.Literal = ident
			tok.Type = token.LookUpIdent(tok.Literal)
			fmt.Printf("scanned %v \n", ident)
			return tok
		} else if isDigit(l.ch) {
			number := l.readNumber()
			tok.Literal = number
			tok.Type = token.INT
			fmt.Printf("scanned %v \n", number)
		} else {
			tok = token.New(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	fmt.Println( l.input[position:l.position])
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
