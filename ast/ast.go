package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
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

// Returns the Literal token of the statement for debug use :D
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // LET token
	Name  *Identifer
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifer struct {
	Token token.Token // token.INDENT
	Value string
}

func (i *Identifer) expressionNode() {}

func (i *Identifer) TokenLiteral() string { return i.Token.Literal }
