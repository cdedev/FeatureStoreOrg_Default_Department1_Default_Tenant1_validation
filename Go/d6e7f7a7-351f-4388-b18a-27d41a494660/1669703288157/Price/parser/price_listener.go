// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669703288157/Price.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Price

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PriceListener is a complete listener for a parse tree produced by PriceParser.
type PriceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrice is called when entering the price production.
	EnterPrice(c *PriceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrice is called when exiting the price production.
	ExitPrice(c *PriceContext)
}
