// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672370801035/Expiry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expiry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExpiryListener is a complete listener for a parse tree produced by ExpiryParser.
type ExpiryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpiry is called when entering the expiry production.
	EnterExpiry(c *ExpiryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpiry is called when exiting the expiry production.
	ExitExpiry(c *ExpiryContext)
}
