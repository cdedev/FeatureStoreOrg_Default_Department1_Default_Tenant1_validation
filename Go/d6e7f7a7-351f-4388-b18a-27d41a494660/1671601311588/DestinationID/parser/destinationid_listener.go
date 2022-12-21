// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671601311588/DestinationID.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DestinationID

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DestinationIDListener is a complete listener for a parse tree produced by DestinationIDParser.
type DestinationIDListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDestinationid is called when entering the destinationid production.
	EnterDestinationid(c *DestinationidContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDestinationid is called when exiting the destinationid production.
	ExitDestinationid(c *DestinationidContext)
}
