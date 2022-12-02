// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972190965/High.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // High

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HighListener is a complete listener for a parse tree produced by HighParser.
type HighListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHigh is called when entering the high production.
	EnterHigh(c *HighContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHigh is called when exiting the high production.
	ExitHigh(c *HighContext)
}
