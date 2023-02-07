// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675788016809/PressureLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PressureLevel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PressureLevelListener is a complete listener for a parse tree produced by PressureLevelParser.
type PressureLevelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPressurelevel is called when entering the pressurelevel production.
	EnterPressurelevel(c *PressurelevelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPressurelevel is called when exiting the pressurelevel production.
	ExitPressurelevel(c *PressurelevelContext)
}
