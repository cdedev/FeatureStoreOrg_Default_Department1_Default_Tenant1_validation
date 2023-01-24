// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674532608460/Margin.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Margin

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MarginListener is a complete listener for a parse tree produced by MarginParser.
type MarginListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMargin is called when entering the margin production.
	EnterMargin(c *MarginContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMargin is called when exiting the margin production.
	ExitMargin(c *MarginContext)
}
