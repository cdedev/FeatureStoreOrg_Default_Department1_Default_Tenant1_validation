// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925607035/Hair.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hair

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HairListener is a complete listener for a parse tree produced by HairParser.
type HairListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHair is called when entering the hair production.
	EnterHair(c *HairContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHair is called when exiting the hair production.
	ExitHair(c *HairContext)
}
