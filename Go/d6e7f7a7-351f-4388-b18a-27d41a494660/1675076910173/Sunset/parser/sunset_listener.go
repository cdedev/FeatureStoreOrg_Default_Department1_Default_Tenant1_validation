// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675076910173/Sunset.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sunset

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SunsetListener is a complete listener for a parse tree produced by SunsetParser.
type SunsetListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSunset is called when entering the sunset production.
	EnterSunset(c *SunsetContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSunset is called when exiting the sunset production.
	ExitSunset(c *SunsetContext)
}
