// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669703494650/OrderDate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OrderDate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OrderDateListener is a complete listener for a parse tree produced by OrderDateParser.
type OrderDateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOrderdate is called when entering the orderdate production.
	EnterOrderdate(c *OrderdateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOrderdate is called when exiting the orderdate production.
	ExitOrderdate(c *OrderdateContext)
}
