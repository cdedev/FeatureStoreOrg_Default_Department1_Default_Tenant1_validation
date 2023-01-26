// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674721161428/RemoteWork.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RemoteWork

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RemoteWorkListener is a complete listener for a parse tree produced by RemoteWorkParser.
type RemoteWorkListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRemotework is called when entering the remotework production.
	EnterRemotework(c *RemoteworkContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRemotework is called when exiting the remotework production.
	ExitRemotework(c *RemoteworkContext)
}
