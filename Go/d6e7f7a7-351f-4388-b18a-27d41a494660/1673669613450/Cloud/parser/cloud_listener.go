// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673669613450/Cloud.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cloud

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CloudListener is a complete listener for a parse tree produced by CloudParser.
type CloudListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCloud is called when entering the cloud production.
	EnterCloud(c *CloudContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCloud is called when exiting the cloud production.
	ExitCloud(c *CloudContext)
}
