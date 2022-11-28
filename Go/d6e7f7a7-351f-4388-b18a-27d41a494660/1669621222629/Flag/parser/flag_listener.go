// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669621222629/Flag.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Flag

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FlagListener is a complete listener for a parse tree produced by FlagParser.
type FlagListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFlag is called when entering the flag production.
	EnterFlag(c *FlagContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFlag is called when exiting the flag production.
	ExitFlag(c *FlagContext)
}
