// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603481435/Actinium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Actinium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ActiniumListener is a complete listener for a parse tree produced by ActiniumParser.
type ActiniumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterActinium is called when entering the actinium production.
	EnterActinium(c *ActiniumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitActinium is called when exiting the actinium production.
	ExitActinium(c *ActiniumContext)
}
