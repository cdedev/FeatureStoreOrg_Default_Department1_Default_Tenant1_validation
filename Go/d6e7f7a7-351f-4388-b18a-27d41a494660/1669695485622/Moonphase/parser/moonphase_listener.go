// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669695485622/Moonphase.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Moonphase

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MoonphaseListener is a complete listener for a parse tree produced by MoonphaseParser.
type MoonphaseListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMoonphase is called when entering the moonphase production.
	EnterMoonphase(c *MoonphaseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMoonphase is called when exiting the moonphase production.
	ExitMoonphase(c *MoonphaseContext)
}
