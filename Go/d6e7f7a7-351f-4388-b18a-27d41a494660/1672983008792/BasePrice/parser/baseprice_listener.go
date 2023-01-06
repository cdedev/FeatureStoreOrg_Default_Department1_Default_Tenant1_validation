// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983008792/BasePrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BasePrice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePriceListener is a complete listener for a parse tree produced by BasePriceParser.
type BasePriceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBaseprice is called when entering the baseprice production.
	EnterBaseprice(c *BasepriceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBaseprice is called when exiting the baseprice production.
	ExitBaseprice(c *BasepriceContext)
}
