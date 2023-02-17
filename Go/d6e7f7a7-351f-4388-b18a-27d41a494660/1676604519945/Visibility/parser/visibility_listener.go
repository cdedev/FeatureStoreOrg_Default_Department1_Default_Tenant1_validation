// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676604519945/Visibility.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Visibility

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VisibilityListener is a complete listener for a parse tree produced by VisibilityParser.
type VisibilityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVisibility is called when entering the visibility production.
	EnterVisibility(c *VisibilityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVisibility is called when exiting the visibility production.
	ExitVisibility(c *VisibilityContext)
}
