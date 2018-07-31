package repl

import (
	"fmt"
	"github.com/miyo/p4_16_compiler_go/lexer"
	"github.com/miyo/p4_16_compiler_go/token"
	"io"
)

func readLexPrintLoop(in io.Reader, out io.Writer) {

	l := lexer.New(in)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}

}

func Start(in io.Reader, out io.Writer) {
	readLexPrintLoop(in, out)
}
