// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673244794403/Order.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Order

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OrderListener is a complete listener for a parse tree produced by OrderParser.
type OrderListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOrder is called when entering the order production.
	EnterOrder(c *OrderContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOrder is called when exiting the order production.
	ExitOrder(c *OrderContext)
}
