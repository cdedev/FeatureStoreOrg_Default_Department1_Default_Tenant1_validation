// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671694797816/Network.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Network

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NetworkListener is a complete listener for a parse tree produced by NetworkParser.
type NetworkListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNetwork is called when entering the network production.
	EnterNetwork(c *NetworkContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNetwork is called when exiting the network production.
	ExitNetwork(c *NetworkContext)
}
