package ast

import (
	"interpretergo/token"
)

type Node interface {
	TokenLiteral() string // returns the literal value of the token it's associated with
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
	if len(p.Statements) > 0 { //The Program node is going to be the root node of every AST our parser produces
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // Name to hold the identifier of the binding
	Value Expression  // Value for the expression that produces the value
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.TokenLiteral() }

type Identifier struct { // implements the Expression interface
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression  // expression that's to be returned
}

// fulfill the Node and Statement interfaces and look identical to the methods defined on *ast.LetStatement
func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.TokenLiteral() }
