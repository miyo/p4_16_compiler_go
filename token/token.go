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

	KEYWORD = "KEYWORD"

	STRING              = "STRING"
	SINGLE_QUOTE_STRING = "SINGLE_QUOTE_STRING"

	COMMENT      = "COMMENT"
	LINE_COMMENT = "LINE_COMMENT"
)

var keywords = map[string]TokenType{
	"action":       ACTION,
	"apply":        APPLY,
	"bit":          KEYWORD,
	"bool":         KEYWORD,
	"const":        CONST,
	"control":      CONTROL,
	"default":      DEFAULT,
	"else":         ELSE,
	"enum":         KEYWORD,
	"error":        KEYWORD,
	"extern":       KEYWORD,
	"exit":         KEYWORD,
	"false":        FALSE,
	"header":       KEYWORD,
	"header_union": KEYWORD,
	"if":           IF,
	"in":           KEYWORD,
	"inout":        KEYWORD,
	"int":          KEYWORD,
	"match_kind":   KEYWORD,
	"package":      KEYWORD,
	"parser":       KEYWORD,
	"out":          OUT,
	"return":       RETURN,
	"select":       KEYWORD,
	"state":        STATE,
	"struct":       KEYWORD,
	"switch":       KEYWORD,
	"table":        TABLE,
	"transition":   KEYWORD,
	"true":         TRUE,
	"tuple":        KEYWORD,
	"typedef":      KEYWORD,
	"varbit":       KEYWORD,
	"verify":       KEYWORD,
	"void":         KEYWORD,
	"metadata":     METADATA,
	"#include":     PREPROCESS,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
