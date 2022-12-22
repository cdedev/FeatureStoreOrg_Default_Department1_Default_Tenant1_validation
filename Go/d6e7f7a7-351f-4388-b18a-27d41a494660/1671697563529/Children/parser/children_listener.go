// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671697563529/Children.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Children

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChildrenListener is a complete listener for a parse tree produced by ChildrenParser.
type ChildrenListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterChildren is called when entering the children production.
	EnterChildren(c *ChildrenContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitChildren is called when exiting the children production.
	ExitChildren(c *ChildrenContext)
}
