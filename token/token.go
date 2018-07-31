package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	SLASH    = "/"
	ASTERISK = "*"
	AND      = "&"
	OR       = "|"

	BANG      = "!"
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	LT = "<"
	GT = ">"

	EQ  = "=="
	NEQ = "!="

	ANDAND  = "&&"
	ORORNEQ = "||"

	BEGIN_COMMENT = "/*"
	END_COMMENT   = "*/"

	// KEYWORDS

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	STRUCT     = "STRUCT"
	HEADER     = "HEADER"
	BIT        = "BIT"
	PREPROCESS = "PREPROCESS"
	CONST      = "CONST"
	STATE      = "STATE"
	DEFAULT    = "DEFAULT"
	CONTROL    = "CONTROL"
	TABLE      = "TABLE"
	ACTION     = "ACTION"
	APPLY      = "APPLY"
	INOUT      = "INOUT"
	OUT        = "OUT"
	METADATA   = "METADATA"

	STRING              = "STRING"
	SINGLE_QUOTE_STRING = "SINGLE_QUOTE_STRING"

	COMMENT      = "COMMENT"
	LINE_COMMENT = "LINE_COMMENT"
)

var keywords = map[string]TokenType{
	"fn":       FUNCTION,
	"let":      LET,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
	"true":     TRUE,
	"false":    FALSE,
	"struct":   FALSE,
	"bit":      FALSE,
	"const":    CONST,
	"state":    STATE,
	"default":  DEFAULT,
	"control":  CONTROL,
	"table":    TABLE,
	"action":   ACTION,
	"apply":    APPLY,
	"inout":    INOUT,
	"out":      OUT,
	"metadata": METADATA,
	"#include": PREPROCESS,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
