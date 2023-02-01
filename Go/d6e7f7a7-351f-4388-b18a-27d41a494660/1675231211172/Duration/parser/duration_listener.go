// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675231211172/Duration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Duration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DurationListener is a complete listener for a parse tree produced by DurationParser.
type DurationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDuration is called when entering the duration production.
	EnterDuration(c *DurationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDuration is called when exiting the duration production.
	ExitDuration(c *DurationContext)
}
