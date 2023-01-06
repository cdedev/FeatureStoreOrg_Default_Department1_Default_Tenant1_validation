// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672983239687/NoOfOrders.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NoOfOrders

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NoOfOrdersListener is a complete listener for a parse tree produced by NoOfOrdersParser.
type NoOfOrdersListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNooforders is called when entering the nooforders production.
	EnterNooforders(c *NoofordersContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNooforders is called when exiting the nooforders production.
	ExitNooforders(c *NoofordersContext)
}
