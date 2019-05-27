package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	BLANK = "BLANK"

	// Identifiers + literals
	IDENT     = "IDENT"   // add, foobar, x, y, ...
	INT       = "INT"     // 1343456
	FLOAT     = "FLOAT"   // 1.23
	STRING    = "STRING"  // "foobar"
	BOOLEAN   = "BOOLEAN" // true or false

	NULL  = "NULL" //
	NIL   = "NIL"
	ERROR = "ERR"

	FINAL = "FINAL"

	// Operators
	ASSIGN   = "="  // implemented
	PLUS     = "+"  // implemented
	MINUS    = "-"  // implemented
	BANG     = "!"  // implemented - factorial not implemented
	ASTERISK = "*"  // implemented
	SLASH    = "/"  // implemented
	POW      = "**" // implemented
	MOD      = "%"  // implemented

	LT  = "<"  // implemented
	GT  = ">"  // implemented
	LTE = "<=" // implemented
	GTE = ">=" // implemented

	AND = "&&"
	OR  = "||"

	EQ     = "==" // implemented
	NOT_EQ = "!=" // implemented

	// Delimiters
	COMMA     = "," // implemented
	SEMICOLON = ";" // implemented - not properly
	COLON     = ":" // implemented

	LPAREN   = "(" // implemented
	RPAREN   = ")" // implemented
	LBRACE   = "{" // implemented
	RBRACE   = "}" // implemented
	LBRACKET = "[" // implemented
	RBRACKET = "]" // implemented

	// Keywords
	FUNCTION = "FUNCTION" // implemented - change!
	FOR      = "FOR"
	WHEN     = "WHEN" // when(x = range arr, step 5) {}
	WHILE    = "WHILE"
	LET      = "LET"   // implemented - not properly
	NEW      = "NEW"   //
	TRUE     = "TRUE"  // implemented
	FALSE    = "FALSE" // implemented
	IF       = "IF"    // implemented
	ELSE     = "ELSE"  // implemented
	SWITCH   = "SWITCH"
	CASE     = "CASE"
	RETURN   = "RETURN" // implemented
	RETURNS  = "RETURNS"
	BREAK    = "BREAK"
	CONTINUE = "CONTINUE"

	// macros
	MACRO = "MACRO" // implemented

	// comments
	OPENCOMMENT      = "<!-~"
	CLOSECOMMENT     = "~-!>"
	MULTILINECOMMENT = "~"

)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"for":      FOR,
	"when":     WHEN,
	"while":    WHILE,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"switch":   SWITCH,
	"case":     CASE,
	"return":   RETURN,
	"macro":    MACRO,
	"break":    BREAK,
	"continue": CONTINUE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
