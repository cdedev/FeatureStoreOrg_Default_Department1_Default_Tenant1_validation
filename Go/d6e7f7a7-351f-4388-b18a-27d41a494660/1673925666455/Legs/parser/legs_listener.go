// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925666455/Legs.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Legs

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LegsListener is a complete listener for a parse tree produced by LegsParser.
type LegsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLegs is called when entering the legs production.
	EnterLegs(c *LegsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLegs is called when exiting the legs production.
	ExitLegs(c *LegsContext)
}
