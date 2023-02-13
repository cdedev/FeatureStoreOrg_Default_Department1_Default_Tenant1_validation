// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676260616947/Length.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Length

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LengthListener is a complete listener for a parse tree produced by LengthParser.
type LengthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLength is called when entering the length production.
	EnterLength(c *LengthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLength is called when exiting the length production.
	ExitLength(c *LengthContext)
}
