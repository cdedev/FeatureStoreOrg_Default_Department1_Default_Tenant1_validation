// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115605743/NumConnections.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NumConnections

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NumConnectionsListener is a complete listener for a parse tree produced by NumConnectionsParser.
type NumConnectionsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNumconnections is called when entering the numconnections production.
	EnterNumconnections(c *NumconnectionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNumconnections is called when exiting the numconnections production.
	ExitNumconnections(c *NumconnectionsContext)
}
