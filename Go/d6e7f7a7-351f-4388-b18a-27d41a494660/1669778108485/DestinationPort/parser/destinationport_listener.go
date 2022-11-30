// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778108485/DestinationPort.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DestinationPort

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DestinationPortListener is a complete listener for a parse tree produced by DestinationPortParser.
type DestinationPortListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDestinationport is called when entering the destinationport production.
	EnterDestinationport(c *DestinationportContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDestinationport is called when exiting the destinationport production.
	ExitDestinationport(c *DestinationportContext)
}
