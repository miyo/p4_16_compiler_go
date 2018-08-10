package ast

import (
	"fmt"
	"p4_16_compiler_go/token"
	"strings"
)

// Expression is interface for expressions
type Expression interface {
	Node
	expressionNode()
}

type Lvalue interface {
	Expression
	lvalue()
}

type IntegerExpr struct {
	Value long
}

func (i *IntegerExpr) expressionNode()      {}
func (i *IntegerExpr) TokenLiteral() string { return "IntegerExpr" }
func (i *IntegerExpr) String() string       { return i.Value.String() }

type BooleanExpr struct {
	Value bool
}

func (b *BooleanExpr) expressionNode()      {}
func (b *BooleanExpr) TokenLiteral() string { return "BooleanExpr" }
func (b *BooleanExpr) String() string       { return b.Value.String() }

type StringExpr struct {
	Value string
}

func (s *StringExpr) expressionNode()      {}
func (s *StringExpr) TokenLiteral() string { return "StringLiteral" }
func (s *StringExpr) String() string       { return s.Value }

type VariableExpr struct {
	Value Expression // nonTypeName: IDENTIFIER, APPLY, KEY, ACTIONS, or STATE
}

func (v *VariableExpr) expressionNode()      {}
func (v *VariableExpr) TokenLiteral() string { return "VariableExpr" }
func (v *VariableExpr) String() string       { return v.String() }

type DotVariableExpr struct { // '.' nonTypeName: IDENTIFIER, APPLY, KEY, ACTIONS, or STATE
	Value Expression // nonTypeName: IDENTIFIER, APPLY, KEY, ACTIONS, or STATE
}

func (v *DotVariableExpr) expressionNode()      {}
func (v *DotVariableExpr) TokenLiteral() string { return "DotVariableExpr" }
func (v *DotVariableExpr) String() string       { return fmt.Sprintf(".%s", v.Value.String()) }

type ArrayAccessExpr struct { // expression '[' expression ']'
	Key   Expression
	Index Expression
}

func (a *ArrayAccessExpr) expressionNode()      {}
func (a *ArrayAccessExpr) TokenLiteral() string { return "ArrayAccessExpr" }
func (a *ArrayAccessExpr) String() string {
	return fmt.SPrintf("%s[%s]", a.Key.String(), a.Index.String())
}

type RangeExpr struct { // expression '[' expression ':' expression ']'
	Key        Expression
	BeginIndex Expression
	EndIndex   Expression
}

func (a *RangeExpr) expressionNode()      {}
func (a *RangeExpr) TokenLiteral() string { return "Range" }
func (a *RangeExpr) String() string {
	return fmt.SPrintf("%s[%s:]", a.Key.String(), a.BeginIndex.String(), a.EndIndex.String())
}

type BlockExpr struct { // '{' expressionList '}'
	Expressions []Expression
}

func (b *BlockExpr) expressionNode()      {}
func (b *BlockExpr) TokenLiteral() string { return "BlockExpr" }
func (b *BlockExpr) String() string {
	s := []string{}
	for _, expr := range b.Expressions {
		s = append(s, expr.String())
	}
	return strings.Join(s, ", ")
}

type ParenExpr struct { // '(' expressionList ')'
	Expr Expression
}

func (p *ParenExpr) expressionNode()      {}
func (p *ParenExpr) TokenLiteral() string { return "ParenExpr" }
func (p *ParenExpr) String() string       { return fmt.Sprintf("(%s)", p.Expr.String()) }

type UnaryExpr struct { // OP expression
	Op   string // '!', '~', '-', or '+'
	Expr Expression
}

func (u *UnaryExpr) expressionNode()      {}
func (u *UnaryExpr) TokenLiteral() string { return "UnaryExpr" }
func (u *UnaryExpr) String() string       { return fmt.Sprintf("%s%s", u.Op, u.Expr.String()) }

type TypeFieldAccessExpr struct { // typeName '.' member
	Type   TypeName
	Member Name
}

func (t *TypeFieldAccessExpr) expressionNode()      {}
func (t *TypeFieldAccessExpr) TokenLiteral() string { return "TypeFieldAccessExpr" }
func (t *TypeFieldAccessExpr) String() string {
	return fmt.Sprintf("%s.%s", t.Type.String(), t.Member.String())
}

type ErrorFieldAccessExpr struct { // ERROR '.' member
	Member Name
}

func (e *ErrorFieldAccessExpr) expressionNode()      {}
func (e *ErrorFieldAccessExpr) TokenLiteral() string { return "ErrorFieldAccessExpr" }
func (e *ErrorFieldAccessExpr) String() string       { return fmt.Sprintf("ERROR.%s", e.Member.String()) }

type InstanceFieldAccessExpr struct { // ERROR '.' member
	Expr   Expression
	Member Name
}

func (i *InstanceFieldAccessExpr) expressionNode()      {}
func (i *InstanceFieldAccessExpr) TokenLiteral() string { return "InstanceFieldAccessExpr" }
func (i *InstanceFieldAccessExpr) String() string {
	return fmt.Sprintf("%s.%s", i.Expr.String(), i.Member.String())
}

type BinaryExpr struct { // expression Op expression
	Op    string // '*', '/', '%', '+', '-', "<<", ">>", "<=", ">=", '<', '>', "!=", "==", '&', '^', '|', "++", "&&", "||"
	Left  Expression
	Right Expression
}

func (b *BinaryExpr) expressionNode()      {}
func (b *BinaryExpr) TokenLiteral() string { return "BinaryExpr" }
func (b *BinaryExpr) String() string {
	return fmt.Sprintf("%s %s %s", b.Left.String(), b.Op, b.Right.String())
}

type ConditionExpr struct { // expression '?' expression ':' expression
	Cond     Expr
	ThenExpr Expression
	ElseExpr Expression
}

func (c *CondExpr) expressionNode()      {}
func (c *CondExpr) TokenLiteral() string { return "CondExpr" }
func (c *CondExpr) String() string {
	return fmt.Sprintf("%s ? %s : %s", c.Cond.String(), c.ThenExpr.String(), c.ElseExpr.String())
}

type CallWithTypeExpr struct { // expression '<' typeArgumentList '>' '(' argumentList ')'
	Key      Expr
	TypeArgs []TypeArg // DONTCARE, typeRef
	Args     []Expression
}

func (c *CallWithTypeExpr) expressionNode()      {}
func (c *CallWithTypeExpr) TokenLiteral() string { return "CallWithTypeExpr" }
func (c *CallWithTypeExpr) String() string {
	ts := []string{}
	for _, t := range c.TypeArgs {
		ts = append(ts, t.String())
	}
	as := []string{}
	for _, a := range c.Args {
		as = append(as, a.String())
	}
	return fmt.Sprintf("%s < %s > ( %s )", c.Key.String(), strings.Join(ts, ","), strings.Join(as, ","))
}

type CallExpr struct { // expression '(' argumentList ')'
	Key  Expr
	Args []Expression
}

func (c *CallExpr) expressionNode()      {}
func (c *CallExpr) TokenLiteral() string { return "CallExpr" }
func (c *CallExpr) String() string {
	as := []string{}
	for _, a := range c.Args {
		as = append(as, a.String())
	}
	return fmt.Sprintf("%s ( %s )", c.Key.String(), strings.Join(as, ","))
}

type TypeExpr struct { // typeRef '(' argumentList ')'
	TypeRef Type
	Args    []Expression
}

func (t *TypeExpr) expressionNode()      {}
func (t *TypeExpr) TokenLiteral() string { return "TypeExpr" }
func (t *TypeExpr) String() string {
	as := []string{}
	for _, a := range c.Args {
		as = append(as, a.String())
	}
	return fmt.Sprintf("%s ( %s )", c.TypeRef.String(), strings.Join(as, ","))
}

type CastExpr struct { // '(' typeRef ')' expression
	TypeRef Type
	Expr    Expression
}

func (c *CastExpr) expressionNode()      {}
func (c *CastExpr) TokenLiteral() string { return "CastExpr" }
func (c *CastExpr) String() string       { return fmt.Sprintf("(%s)%s", c.TypeRef.String(), c.Expr.String()) }

type LvaluePrefixedNonTypeName struct {
	Values []Expression // nonTypeName: IDENTIFIER, APPLY, KEY, ACTIONS, or STATE
}

func (l *LvaluePrefixedNonTypeName) lvalue()              {}
func (l *LvaluePrefixedNonTypeName) expressionNode()      {}
func (l *LvaluePrefixedNonTypeName) TokenLiteral() string { return "LvaluePrefixedNonTypeName" }
func (l *LvaluePrefixedNonTypeName) String() string {
	s := []string{}
	for _, v := range Values {
		s = append(s, v.String())
	}
	return string.Join(s, ".")
}

type LvalueWithMember struct {
	Value  Lvalue
	Member Name
}

func (l *LvalueWithMember) lvalue()              {}
func (l *LvalueWithMember) expressionNode()      {}
func (l *LvalueWithMember) TokenLiteral() string { return "LvalueWithMember" }
func (l *LvalueWithMember) String() string {
	return fmt.Sprintf("%s.%s", l.Value.String(), l.Member.String())
}

type LvalueWithIndex struct {
	Value Lvalue
	Index Expression
}

func (l *LvalueWithIndex) lvalue()              {}
func (l *LvalueWithIndex) expressionNode()      {}
func (l *LvalueWithIndex) TokenLiteral() string { return "LvalueWithIndex" }
func (l *LvalueWithIndex) String() string {
	return fmt.Sprintf("%s[%s]", l.Value.String(), l.Index.String())
}

type LvalueWithRange struct {
	Value      Lvalue
	BeginIndex Expression
	EndIndex   Expression
}

func (l *LvalueWithRange) lvalue()              {}
func (l *LvalueWithRange) expressionNode()      {}
func (l *LvalueWithRange) TokenLiteral() string { return "LvalueWithRange" }
func (l *LvalueWithRange) String() string {
	return fmt.Sprintf("%s[%s:%s]", l.Value.String(), l.BeginIndex.String(), l.EndIndex.String())
}
