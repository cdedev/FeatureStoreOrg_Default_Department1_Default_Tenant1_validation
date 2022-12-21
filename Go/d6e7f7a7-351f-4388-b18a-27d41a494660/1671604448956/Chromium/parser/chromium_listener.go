// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604448956/Chromium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chromium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChromiumListener is a complete listener for a parse tree produced by ChromiumParser.
type ChromiumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterChromium is called when entering the chromium production.
	EnterChromium(c *ChromiumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitChromium is called when exiting the chromium production.
	ExitChromium(c *ChromiumContext)
}
