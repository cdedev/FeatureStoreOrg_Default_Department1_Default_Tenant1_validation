// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656699912/DewPoint.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DewPoint

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DewPointListener is a complete listener for a parse tree produced by DewPointParser.
type DewPointListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDewpoint is called when entering the dewpoint production.
	EnterDewpoint(c *DewpointContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDewpoint is called when exiting the dewpoint production.
	ExitDewpoint(c *DewpointContext)
}
