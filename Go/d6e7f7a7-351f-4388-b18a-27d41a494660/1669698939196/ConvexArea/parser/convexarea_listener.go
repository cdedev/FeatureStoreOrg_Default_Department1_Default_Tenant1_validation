// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669698939196/ConvexArea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ConvexArea

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConvexAreaListener is a complete listener for a parse tree produced by ConvexAreaParser.
type ConvexAreaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConvexarea is called when entering the convexarea production.
	EnterConvexarea(c *ConvexareaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConvexarea is called when exiting the convexarea production.
	ExitConvexarea(c *ConvexareaContext)
}
