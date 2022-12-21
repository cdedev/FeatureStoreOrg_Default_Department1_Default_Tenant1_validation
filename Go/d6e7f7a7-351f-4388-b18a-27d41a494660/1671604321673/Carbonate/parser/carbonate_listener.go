// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604321673/Carbonate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbonate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CarbonateListener is a complete listener for a parse tree produced by CarbonateParser.
type CarbonateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCarbonate is called when entering the carbonate production.
	EnterCarbonate(c *CarbonateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCarbonate is called when exiting the carbonate production.
	ExitCarbonate(c *CarbonateContext)
}
