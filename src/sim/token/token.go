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
	STRUCT    = "STRUCT"
	INTERFACE = "INTERFACE"
	ENUM      = "ENUM"
	CAST      = "CAST"
	VOID      = "VOID"
	STATIC    = "STATIC" // in this case, it will be used for constants		unimplemented 21st
	// unimplemented - casting
	// TODO: Primitives

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

	INCREMENT = "++"
	DECREMENT = "--"
	PLUSEQ    = "+="
	MINUSEQ   = "-="
	MULTEQ    = "*="
	DIVEQ     = "/="
	MODEQ     = "%="
	POWEQ     = "**="

	// bitwise
	XOREQ  = "^="
	OREQ   = "|="
	ANDEQ  = "&="
	RSHIFT = ">>"
	LSHIFT = "<<"

	BOR  = "|"
	BAND = "&"
	XOR  = "^"

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
	PACKAGE  = "PACKAGE" //
	IMPORT   = "USE"
	FROM     = "FROM"     //
	FUNCTION = "FUNCTION" // implemented - change!
	METHOD   = "METHOD"
	DEFINE   = "DEF" //
	CLASS    = "CLASS"
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
	HOST     = "HOST" // host(port; file; param) - e.g. host(4242; index.html; username, message -> {action})
	EXTENDS  = "EXTENDS"
	TYPE     = "TYPE" //
	THIS     = "THIS"
	OVERRIDE = "OVERRIDE" //
	TRY      = "TRY"      //
	CATCH    = "CATCH"    //
	SUPER    = "SUPER"
	THROW    = "THROW"   //
	THROWS   = "THROWS"  //
	DECLARE  = "DECLARE" //

	// macros
	MACRO = "MACRO" // implemented

	// comments
	OPENCOMMENT      = "<!-~"
	CLOSECOMMENT     = "~-!>"
	MULTILINECOMMENT = "~"
	// <!-~
	// ~
	// ~
	// ~-!>

	// access modifiers
	PUBLIC  = "GLOBAL"
	PRIVATE = "LOCAL"

	// pointers
	AT = "@" // points to object

	AS = "AS" // type declaration: let x be 5 as string. declare y as string. - can be used for casting, let x be 5. let z be x as string
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"package":  PACKAGE,
	"use":      IMPORT,
	"from":     FROM,
	"fn":       FUNCTION,
	"method":   METHOD,
	"def":      DEFINE,
	"class":    CLASS,
	"for":      FOR,
	"when":     WHEN,
	"while":    WHILE,
	"let":      LET,
	"new":      NEW,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"switch":   SWITCH,
	"case":     CASE,
	"return":   RETURN,
	"returns":  RETURNS,
	"macro":    MACRO,
	"global":   PUBLIC,
	"local":    PRIVATE,
	"break":    BREAK,
	"continue": CONTINUE,
	"HOST":     HOST,
	"EXTENDS":  EXTENDS,
	"type":     TYPE,
	"this":     THIS,
	"Override": OVERRIDE,
	"try":      TRY,
	"catch":    CATCH,
	"super":    SUPER,
	"throws":   THROWS,
	"throw":    THROW,
	"declare":  DECLARE,
	"As":       AS,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
