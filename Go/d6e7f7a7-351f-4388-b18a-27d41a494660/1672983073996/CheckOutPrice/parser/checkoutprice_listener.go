// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983073996/CheckOutPrice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CheckOutPrice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CheckOutPriceListener is a complete listener for a parse tree produced by CheckOutPriceParser.
type CheckOutPriceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCheckoutprice is called when entering the checkoutprice production.
	EnterCheckoutprice(c *CheckoutpriceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCheckoutprice is called when exiting the checkoutprice production.
	ExitCheckoutprice(c *CheckoutpriceContext)
}
