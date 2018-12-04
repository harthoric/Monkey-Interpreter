package lexer

import (
	"Simplex-Simia/token"
	"fmt"
)

var Count int

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	tokDict := map[string]token.TokenType{
		"=":  token.ASSIGN,
		"==": token.EQ,
		"+":  token.PLUS,
		"-":  token.MINUS,
		"!":  token.BANG,
		"!=": token.NOT_EQ,
		"%":  token.MOD,
		"/":  token.SLASH,
		"*":  token.ASTERISK,
		"**": token.POW,
		"<":  token.LT,
		"<=": token.LTE,
		">":  token.GT,
		">=": token.GTE,
		";":  token.SEMICOLON,
		":":  token.COLON,
		",":  token.COMMA,
		"{":  token.LBRACE,
		"}":  token.RBRACE,
		"(":  token.LPAREN,
		")":  token.RPAREN,
		"[":  token.LBRACKET,
		"]":  token.RBRACKET,
	}

	if val, ok := tokDict[string(l.ch)]; ok {
		if l.peekChar() == '=' || l.peekChar() == '*' {
			oldPCh := l.peekChar()
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: tokDict[string(l.ch)+string(oldPCh)], Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(val, l.ch)
		}
	} else {
		if l.ch == '"' {
			tok.Type = token.STRING
			tok.Literal = l.readString()
		} else if l.ch == 0 {
			tok.Literal = ""
			tok.Type = token.EOF
		} else if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			fmt.Println(string(l.ch))
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' || l.ch == '|' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	Count++
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
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || ch == '.'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
