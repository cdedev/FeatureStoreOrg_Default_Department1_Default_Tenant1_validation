// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118968290/RideLength.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RideLength

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RideLengthListener is a complete listener for a parse tree produced by RideLengthParser.
type RideLengthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRidelength is called when entering the ridelength production.
	EnterRidelength(c *RidelengthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRidelength is called when exiting the ridelength production.
	ExitRidelength(c *RidelengthContext)
}
