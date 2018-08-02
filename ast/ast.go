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

// Expression is interface for expressions
type Expression interface {
	Node
	expressionNode()
}

type Type interface {
	Node
	typeNode()
}

type P4Program struct {
	Declarations []Declaration
}

func (p *P4Program) TokenLiteral() { return "P4Program" }
func (p *P4Program) String()       { return "P4Program" }

type ConstantDeclaration struct {
	Annotations []*Annotation
	TypeRef     Type
	Name        string
	Initializer Expression
}

func (d *ConstantDeclaration) declarationNode() {}
func (d *ConstantDeclaration) TokenLiteral()    { return "Constant" }
func (d *ConstantDeclaration) String()          { return "Constant" }

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

type ExternDeclaration struct {
}

func (d *ExternDeclaration) declarationNode() {}
func (d *ExternDeclaration) TokenLiteral()    { return "Extern" }
func (d *ExternDeclaration) String()          { return "Extern" }

type ActionDeclaration struct {
}

func (d *ActionDeclaration) declarationNode() {}
func (d *ActionDeclaration) TokenLiteral()    { return "Action" }
func (d *ActionDeclaration) String()          { return "Action" }

type ParserDeclaration struct {
}

func (d *ParserDeclaration) declarationNode() {}
func (d *ParserDeclaration) TokenLiteral()    { return "Parser" }
func (d *ParserDeclaration) String()          { return "Parser" }

type TypeDeclaration struct {
}

func (d *TypeDeclaration) declarationNode() {}
func (d *TypeDeclaration) TokenLiteral()    { return "Type" }
func (d *TypeDeclaration) String()          { return "Type" }

type ControlDeclaration struct {
}

func (d *ControlDeclaration) declarationNode() {}
func (d *ControlDeclaration) TokenLiteral()    { return "Control" }
func (d *ContorlDeclaration) String()          { return "Control" }

type Instantiation struct {
}

func (d *Instantiation) declarationNode() {}
func (d *Instantiation) TokenLiteral()    { return "Instantiation" }
func (d *Instantiation) String()          { return "Instantiation" }

type ErrorDeclaration struct {
}

func (d *ErrorDeclaration) declarationNode() {}
func (d *ErrorDeclaration) TokenLiteral()    { return "Error" }
func (d *ErrorDeclaration) String()          { return "Error" }

type MatchKindDeclaration struct {
}

func (d *MatchKindDeclaration) declarationNode() {}
func (d *MatchKindDeclaration) TokenLiteral()    { return "MatchKind" }
func (d *MatchKindDeclaration) String()          { return "MatchKind" }
