// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674534015928/Rural.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rural

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RuralListener is a complete listener for a parse tree produced by RuralParser.
type RuralListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRural is called when entering the rural production.
	EnterRural(c *RuralContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRural is called when exiting the rural production.
	ExitRural(c *RuralContext)
}
