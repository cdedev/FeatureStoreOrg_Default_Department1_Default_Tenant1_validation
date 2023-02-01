// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675242465648/Discount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Discount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiscountListener is a complete listener for a parse tree produced by DiscountParser.
type DiscountListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDiscount is called when entering the discount production.
	EnterDiscount(c *DiscountContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDiscount is called when exiting the discount production.
	ExitDiscount(c *DiscountContext)
}
