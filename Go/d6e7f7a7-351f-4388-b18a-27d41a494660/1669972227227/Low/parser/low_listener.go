// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972227227/Low.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Low

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LowListener is a complete listener for a parse tree produced by LowParser.
type LowListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLow is called when entering the low production.
	EnterLow(c *LowContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLow is called when exiting the low production.
	ExitLow(c *LowContext)
}
