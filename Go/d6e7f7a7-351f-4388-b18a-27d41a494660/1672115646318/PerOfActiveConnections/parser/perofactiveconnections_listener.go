// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115646318/PerOfActiveConnections.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PerOfActiveConnections

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PerOfActiveConnectionsListener is a complete listener for a parse tree produced by PerOfActiveConnectionsParser.
type PerOfActiveConnectionsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPerofactiveconnections is called when entering the perofactiveconnections production.
	EnterPerofactiveconnections(c *PerofactiveconnectionsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPerofactiveconnections is called when exiting the perofactiveconnections production.
	ExitPerofactiveconnections(c *PerofactiveconnectionsContext)
}
