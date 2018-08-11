package ast

import (
	"buffers"
	"fmt"
	"p4_16_compiler_go/token"
	"strings"
)

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

type TableDeclaration struct {
	// : optAnnotations TABLE name '{' tablePropertyList '}'
	Annotations     []Annotation
	Name            Name
	TableProperties TableProperty
}

type TableProperty interface {
	Node
	tablePorperty()
}

type TablePropertyOfKeyList struct {
	// : KEY '=' '{' keyElementList '}'
	KeyElements []KeyElement
}

func (t *TablePropertyOfKeyList) tableProperty()       {}
func (t *TablePropertyOfKeyList) TokenLiteral() string { "TablePropertyOfKeyList" }
func (t *TablePropertyOfKeyList) String() string {
	s := []string{}
	for _, e := range KeyElements {
		s = append(s, e.String())
	}
	return fmt.Sprintf("key = %s", s)
}

type TablePropertyOfActionList struct {
	// | ACTIONS '=' '{' actionList '}'
	ActionRefs []AcctionRef
}

type TablePropertyOfImmutableEntry struct {
	// | CONST ENTRIES '=' '{' entriesList '}' /* immutable entries */
	Entries []Entry
}

type TablePropertyOfIdentifier struct {
	// | optAnnotations CONST IDENTIFIER '=' initializer ';'
	// | optAnnotations IDENTIFIER '=' initializer ';'
	Annotations []Annotation
	Const       bool
	Ident       Identifier
	Initializer Expression
}

type keyElement struct {
	// : expression ':' name optAnnotations ';'
	Expr        Expression
	Name        Name
	Annotations []Annotation
}

type Entry struct {
	// : optAnnotations keysetExpression ':' actionRef ';'
	Annotations []Annotation
	KeysetExpr  KeySetExpression
	ActionRef   AcctionRef
}

type ActionRef struct {
	// : optAnnotations name
	// | optAnnotations name '(' argumentList ')'
	Annotations []Annotation
	Name        Name
	Args        []Expression
}

type ActionDeclaration struct {
	// : optAnnotations ACTION name '(' parameterList ')' blockStatement
	Annotations []Annotation
	Name        Name
	Parameters  []Parameter
	Body        BlockStatement
}

func (d *ActionDeclaration) declarationNode() {}
func (d *ActionDeclaration) TokenLiteral()    { return "Action" }
func (d *ActionDeclaration) String() {
	var out buffers.ByteBuffer
	{
		s := []string{}
		for _, a := range d.Annotations {
			s = append(s, a.String())
		}
		fmt.Fprintf(out, "%s ", string.Join(s, ", "))
	}
	fmt.Fprintf(out, "action %s ", d.Name)
	{
		s := []string{}
		for _, p := range d.Parameters {
			s = append(s, p.String())
		}
		fmt.Fprintf(out, "( %s )", string.Join(s, ", "))
	}
	fmt.Fprintf(out, "%s", d.Body.String())
	return out.String()
}

type VariableDeclaration struct {
	// : annotations typeRef name optInitializer ';'
	// | typeRef name optInitializer ';'
	Annotations []Annotation
	Type        TypeRef
	Name        Name
	Initializer Expression
}

func (d *VariableDeclaration) statementNode() {}
func (d *VariableDeclaration) TokenLiteral()  { return "VariableDeclaration" }
func (d *VarialbeDeclaration) String() {
	var out buffers.ByteBuffer
	s := []string{}
	for _, a := range d.Annotations {
		s = append(s, a.String())
	}
	fmt.Fprintf(out, "%s %s %s", string.Join(s, ", "), d.Type.String(), d.Type.String())
	if d.Initializer != nil {
		fmt.Fprintf(out, " = %s", d.Initializer.String())
	}
	fmt.Fprintf(out, ";")
	return out.String()
}

type ConstantDeclaration struct {
	// : optAnnotations CONST typeRef name '=' initializer ';'
	Annotations []Annotation
	Type        TypeRef
	Name        Name
	Initializer Expression
}

func (d *ConstantDeclaration) statementNode() {}
func (d *ConstantDeclaration) TokenLiteral()  { return "ConstantDeclaration" }
func (d *VarialbeDeclaration) String() {
	var out buffers.ByteBuffer
	s := []string{}
	for _, a := range d.Annotations {
		s = append(s, a.String())
	}
	fmt.Fprintf(out, "%s const %s %s = ;", string.Join(s, ", "), d.Type.String(), d.Type.String(), d.Initializer.String())
	return out.String()
}
