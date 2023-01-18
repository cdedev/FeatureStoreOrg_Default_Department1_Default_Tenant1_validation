// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674015061824/Monetary.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Monetary

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MonetaryListener is a complete listener for a parse tree produced by MonetaryParser.
type MonetaryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMonetary is called when entering the monetary production.
	EnterMonetary(c *MonetaryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMonetary is called when exiting the monetary production.
	ExitMonetary(c *MonetaryContext)
}
