package ast

import (
	"p4_16_compiler_go/token"
)

type ParserDeclaration struct {
	Type ParserTypeDeclaration
	Body []*ParserDeclBlock
}

func (d *ParserDeclaration) declarationNode() {}
func (d *ParserDeclaration) TokenLiteral()    { return "Parser" }
func (d *ParserDeclaration) String()          { return "Parser" }

type ParserDeclBlock struct {
	ParserLocalElements []Declaration // constantDeclaration, variableDeclaration, or instantiation
	ParseStates         []*ParserState
}

type ParserState struct {
	Annotations          []*Annotation
	Name                 Name
	ParserStatements     []ParserStatement
	TransitionStatements []*TransitionStatement
}

type ParserStatement interface {
	parserStatement()
}

type TransitionStatement struct {
	StateExpr StateExpression
}

type StateExpression struct {
	Name
	SelectExpr SelectExpression
}

type SelectExpression struct {
	Expressions []Exression
	SelectCases []SelectCase
}

type SelectCase struct {
	KyesetExpression Expression
	Name             string
}

type TupleKeysetExpression struct {
	SimpleKeysetExpr SimpleKeysetExpression
}

func (*TupleKeysetExpression) expressionNode() {}
