// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671526892617/Liquid.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Liquid

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LiquidListener is a complete listener for a parse tree produced by LiquidParser.
type LiquidListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLiquid is called when entering the liquid production.
	EnterLiquid(c *LiquidContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLiquid is called when exiting the liquid production.
	ExitLiquid(c *LiquidContext)
}
