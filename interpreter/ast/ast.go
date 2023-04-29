package ast

import (
	"bytes"
	"strings"

	"github.com/wule61/derive/interpreter/token"
)

type Node interface {
	// TokenLiteral() 仅用于调试和测试
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type ExpressionStatement struct {
	Token      token.Token // 该表达式中第一个词法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es != nil {
		return es.Expression.String()
	}
	return ""
}

type PrefixExpression struct {
	Token    token.Token // 前缀词法单元 比如：! -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode()      {}
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }
func (b *Boolean) String() string       { return b.Token.Literal }

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	var elements []string
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString(token.LBRACKET)
	out.WriteString(strings.Join(elements, token.COMMA+" "))
	out.WriteString(token.RBRACKET)

	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}
func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString(token.LPAREN)
	out.WriteString(strings.Join(args, token.COMMA+" "))
	out.WriteString(token.RPAREN)

	return out.String()
}

type ShapStatement struct {
	Token    token.Token // the token.SHAP token
	Elements []Expression
}

func (ss *ShapStatement) statementNode() {}
func (ss *ShapStatement) TokenLiteral() string {
	return ss.Token.Literal
}

func (ss *ShapStatement) String() string {
	var out bytes.Buffer

	var args []string
	for _, a := range ss.Elements {
		args = append(args, a.String())
	}

	out.WriteString(ss.Token.Literal)
	out.WriteString(token.LBRACKET)
	out.WriteString(strings.Join(args, token.COMMA+" "))
	out.WriteString(token.RBRACKET)

	return out.String()
}

type AssignLiteral struct {
	Token      token.Token // The '=' token
	Name       Expression
	Expression Expression
}

func (al *AssignLiteral) expressionNode() {}
func (al *AssignLiteral) TokenLiteral() string {
	return al.Token.Literal
}
func (al *AssignLiteral) String() string {
	var out bytes.Buffer

	out.WriteString(al.Name.String() + " ")
	out.WriteString(al.Token.Literal + " ")
	out.WriteString(al.Expression.String())
	return out.String()
}
