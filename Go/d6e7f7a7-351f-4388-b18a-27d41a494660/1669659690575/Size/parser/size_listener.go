// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669659690575/Size.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Size

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SizeListener is a complete listener for a parse tree produced by SizeParser.
type SizeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSize is called when entering the size production.
	EnterSize(c *SizeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSize is called when exiting the size production.
	ExitSize(c *SizeContext)
}
