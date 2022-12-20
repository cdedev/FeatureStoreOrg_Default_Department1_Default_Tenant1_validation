// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671517555286/Brightness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Brightness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BrightnessListener is a complete listener for a parse tree produced by BrightnessParser.
type BrightnessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBrightness is called when entering the brightness production.
	EnterBrightness(c *BrightnessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBrightness is called when exiting the brightness production.
	ExitBrightness(c *BrightnessContext)
}
