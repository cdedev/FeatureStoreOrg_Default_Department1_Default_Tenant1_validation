// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674794922406/Service.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Service

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ServiceListener is a complete listener for a parse tree produced by ServiceParser.
type ServiceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterService is called when entering the service production.
	EnterService(c *ServiceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitService is called when exiting the service production.
	ExitService(c *ServiceContext)
}
