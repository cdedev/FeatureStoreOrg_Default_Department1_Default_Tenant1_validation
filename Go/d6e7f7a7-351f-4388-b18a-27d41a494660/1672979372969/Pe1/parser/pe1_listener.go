// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979372969/Pe1.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pe1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Pe1Listener is a complete listener for a parse tree produced by Pe1Parser.
type Pe1Listener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPe1 is called when entering the pe1 production.
	EnterPe1(c *Pe1Context)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPe1 is called when exiting the pe1 production.
	ExitPe1(c *Pe1Context)
}
