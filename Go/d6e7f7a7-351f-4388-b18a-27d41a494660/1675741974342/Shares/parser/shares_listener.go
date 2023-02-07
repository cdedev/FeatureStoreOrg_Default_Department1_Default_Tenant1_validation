// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675741974342/Shares.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shares

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SharesListener is a complete listener for a parse tree produced by SharesParser.
type SharesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterShares is called when entering the shares production.
	EnterShares(c *SharesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitShares is called when exiting the shares production.
	ExitShares(c *SharesContext)
}
