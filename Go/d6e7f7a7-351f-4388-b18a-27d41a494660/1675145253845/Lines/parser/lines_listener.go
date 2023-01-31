// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675145253845/Lines.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lines

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LinesListener is a complete listener for a parse tree produced by LinesParser.
type LinesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLines is called when entering the lines production.
	EnterLines(c *LinesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLines is called when exiting the lines production.
	ExitLines(c *LinesContext)
}
