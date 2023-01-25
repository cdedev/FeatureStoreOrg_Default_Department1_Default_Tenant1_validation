// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674667920330/Cash.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cash

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CashListener is a complete listener for a parse tree produced by CashParser.
type CashListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCash is called when entering the cash production.
	EnterCash(c *CashContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCash is called when exiting the cash production.
	ExitCash(c *CashContext)
}
