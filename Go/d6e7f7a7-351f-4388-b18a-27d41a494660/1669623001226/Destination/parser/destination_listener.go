// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623001226/Destination.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Destination

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DestinationListener is a complete listener for a parse tree produced by DestinationParser.
type DestinationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDestination is called when entering the destination production.
	EnterDestination(c *DestinationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDestination is called when exiting the destination production.
	ExitDestination(c *DestinationContext)
}
