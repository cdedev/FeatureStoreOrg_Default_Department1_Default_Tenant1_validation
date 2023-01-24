// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674580219092/Rain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rain

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RainListener is a complete listener for a parse tree produced by RainParser.
type RainListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRain is called when entering the rain production.
	EnterRain(c *RainContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRain is called when exiting the rain production.
	ExitRain(c *RainContext)
}
