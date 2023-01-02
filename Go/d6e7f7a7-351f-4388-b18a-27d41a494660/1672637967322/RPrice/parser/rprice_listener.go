// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637967322/RPrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RPrice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RPriceListener is a complete listener for a parse tree produced by RPriceParser.
type RPriceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRprice is called when entering the rprice production.
	EnterRprice(c *RpriceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRprice is called when exiting the rprice production.
	ExitRprice(c *RpriceContext)
}
