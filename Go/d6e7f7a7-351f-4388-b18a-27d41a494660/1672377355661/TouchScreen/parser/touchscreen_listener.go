// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377355661/TouchScreen.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TouchScreen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TouchScreenListener is a complete listener for a parse tree produced by TouchScreenParser.
type TouchScreenListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTouchscreen is called when entering the touchscreen production.
	EnterTouchscreen(c *TouchscreenContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTouchscreen is called when exiting the touchscreen production.
	ExitTouchscreen(c *TouchscreenContext)
}
