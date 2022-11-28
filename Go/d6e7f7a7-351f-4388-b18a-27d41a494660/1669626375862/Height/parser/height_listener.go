// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626375862/Height.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Height

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HeightListener is a complete listener for a parse tree produced by HeightParser.
type HeightListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHeight is called when entering the height production.
	EnterHeight(c *HeightContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHeight is called when exiting the height production.
	ExitHeight(c *HeightContext)
}
