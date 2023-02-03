// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675404923351/Tilt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tilt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TiltListener is a complete listener for a parse tree produced by TiltParser.
type TiltListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTilt is called when entering the tilt production.
	EnterTilt(c *TiltContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTilt is called when exiting the tilt production.
	ExitTilt(c *TiltContext)
}
