package ast

import (
	"p4_16_compiler_go/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is interface for statements
type Declaration interface {
	Node
	declarationNode()
}

type Type interface {
	Node
	typeNode()
}

type NonTypeName interface {
	Name
	nonTypeName()
}

type Name interface {
	name()
}

type Direction int

const (
	IN Direction = iota
	OUT
	INOUT
)

type P4Program struct {
	Declarations []Declaration
}

func (p *P4Program) TokenLiteral() { return "P4Program" }
func (p *P4Program) String()       { return "P4Program" }

type Annotation struct {
	Name  Identifier
	Exprs []Expression
}

type Identifier struct {
	Token Token.Literal
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral()   { return i.Token.Literal }
func (i *Identifier) expressionNode() { return i.Value }
