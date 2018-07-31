package lexer

import (
	"bufio"
	"bytes"
	"github.com/miyo/p4_16_compiler_go/token"
	"io"
)

type Lexer struct {
	input *bufio.Reader
	ch    byte
	buf   []byte
	out   bytes.Buffer
}

type LexState int

const (
	Normal LexState = iota
	InComment
)

func New(input io.Reader) *Lexer {
	l := &Lexer{input: bufio.NewReaderSize(input, 1), buf: make([]byte, 1)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	_, err := l.input.Read(l.buf)
	if err == io.EOF {
		l.ch = 0
	} else {
		l.ch = l.buf[0]
		l.out.Write(l.buf)
	}
}

func (l *Lexer) peekChar() byte {
	b, err := l.input.Peek(1)
	if err == io.EOF {
		return 0
	} else {
		return b[0]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case '"', '\'':
		if l.ch == '"' {
			tok.Type = token.STRING
		} else {
			tok.Type = token.SINGLE_QUOTE_STRING
		}
		tok.Literal = l.readString(l.ch)
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NEQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '#':
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	case '/':
		if l.peekChar() == '*' {
			l.readChar()
			tok.Type = token.COMMENT
			tok.Literal = l.readComment()
		} else if l.peekChar() == '/' {
			l.readChar()
			tok.Type = token.LINE_COMMENT
			tok.Literal = l.readLineComment()
		} else {
			tok = newToken(token.SLASH, l.ch)
		}
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.ANDAND, Literal: literal}
		} else {
			tok = newToken(token.AND, l.ch)
		}
	case '|':
		tok = newToken(token.OR, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.out.Reset()
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	for {
		ch := l.peekChar()
		if !(isLetter(ch) || isDigit(ch)) {
			break
		}
		l.readChar()
	}
	defer l.readChar()
	defer l.out.Reset()
	return l.out.String()
}

func (l *Lexer) readNumber() string {
	for isDigit(l.peekChar()) {
		l.readChar()
	}
	defer l.readChar()
	defer l.out.Reset()
	return l.out.String()
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '.'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.out.Reset()
		l.readChar()
	}
}

func (l *Lexer) readString(ch byte) string {
	l.out.Reset()
	for {
		l.readChar()
		if l.peekChar() == ch || l.peekChar() == 0 {
			break
		}
	}
	defer l.readChar()
	defer l.out.Reset()
	return l.out.String()
}

func (l *Lexer) readComment() string {
	for {
		l.readChar()
		if l.ch == 0 {
			break
		} else if l.ch == '*' && l.peekChar() == '/' {
			l.readChar()
			break
		}
	}
	defer l.out.Reset()
	s := l.out.String()
	return s[2 : len(s)-2]
}

func (l *Lexer) readLineComment() string {
	for {
		l.readChar()
		if l.ch == 0 {
			break
		} else if l.ch == '\n' {
			break
		}
	}
	defer l.out.Reset()
	s := l.out.String()
	return s[2 : len(s)-1]
}
