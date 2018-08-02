package ast

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

type P4Program struct {
	Declarations []Declaration
}

func (p *P4Program) TokenLiteral() { return "P4Program" }
func (p *P4Program) String()       { return "P4Program" }

type ConstantDeclaration struct {
}

func (d *ConstantDeclaration) declarationNode() {}
func (d *ConstantDeclaration) TokenLiteral()    { return "Constant" }

type ExternDeclaration struct {
}
type ActionDeclaration struct {
}
type ParserDeclaration struct {
}
type TypeDeclaration struct {
}
type ControlDeclaration struct {
}
type Instantiation struct {
}
type ErrorDeclaration struct {
}
type matchKindDeclaration struct {
}
