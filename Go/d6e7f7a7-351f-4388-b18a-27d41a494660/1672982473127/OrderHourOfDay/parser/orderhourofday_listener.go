// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672982473127/OrderHourOfDay.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OrderHourOfDay

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OrderHourOfDayListener is a complete listener for a parse tree produced by OrderHourOfDayParser.
type OrderHourOfDayListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOrderhourofday is called when entering the orderhourofday production.
	EnterOrderhourofday(c *OrderhourofdayContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOrderhourofday is called when exiting the orderhourofday production.
	ExitOrderhourofday(c *OrderhourofdayContext)
}
