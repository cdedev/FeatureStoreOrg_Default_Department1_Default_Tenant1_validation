// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604495022/Cobalt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cobalt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CobaltListener is a complete listener for a parse tree produced by CobaltParser.
type CobaltListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCobalt is called when entering the cobalt production.
	EnterCobalt(c *CobaltContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCobalt is called when exiting the cobalt production.
	ExitCobalt(c *CobaltContext)
}
