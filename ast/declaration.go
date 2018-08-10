package ast

import (
	"p4_16_compiler_go/token"
)

// Statement is interface for statements
type Declaration interface {
	Node
	declarationNode()
}

type ConstantDeclaration struct {
	Annotations []*Annotation
	TypeRef     Type
	Name        string
	Initializer Expression
}

func (d *ConstantDeclaration) declarationNode() {}
func (d *ConstantDeclaration) TokenLiteral()    { return "Constant" }
func (d *ConstantDeclaration) String()          { return "Constant" }

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
