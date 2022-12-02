// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972069640/AdjClose.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AdjClose

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AdjCloseListener is a complete listener for a parse tree produced by AdjCloseParser.
type AdjCloseListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAdjclose is called when entering the adjclose production.
	EnterAdjclose(c *AdjcloseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAdjclose is called when exiting the adjclose production.
	ExitAdjclose(c *AdjcloseContext)
}
