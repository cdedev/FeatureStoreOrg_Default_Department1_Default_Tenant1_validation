// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669701718101/Shipping.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Shipping

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ShippingListener is a complete listener for a parse tree produced by ShippingParser.
type ShippingListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterShipping is called when entering the shipping production.
	EnterShipping(c *ShippingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitShipping is called when exiting the shipping production.
	ExitShipping(c *ShippingContext)
}
