// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675318303202/Wind.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wind

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WindListener is a complete listener for a parse tree produced by WindParser.
type WindListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWind is called when entering the wind production.
	EnterWind(c *WindContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWind is called when exiting the wind production.
	ExitWind(c *WindContext)
}
