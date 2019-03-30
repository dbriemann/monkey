package ast

import "github.com/dbriemann/monkey/token"

// Node represents an abstract node in the syntax tree.
type Node interface {
	TokenLiteral() string
}

// Statement is a statement node.
type Statement interface {
	Node
	statementNode()
}

// Expression is an expression node.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of an AST.
type Program struct {
	Statements []Statement
}

// TokenLiteral satisfies the Node interface.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
