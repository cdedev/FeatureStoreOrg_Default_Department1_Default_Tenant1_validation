// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669635873840/Delivery_Time.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Delivery_Time

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Delivery_TimeListener is a complete listener for a parse tree produced by Delivery_TimeParser.
type Delivery_TimeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDelivery_time is called when entering the delivery_time production.
	EnterDelivery_time(c *Delivery_timeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDelivery_time is called when exiting the delivery_time production.
	ExitDelivery_time(c *Delivery_timeContext)
}
