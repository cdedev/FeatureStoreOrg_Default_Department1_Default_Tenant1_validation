// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669652405435/Failure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Failure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FailureListener is a complete listener for a parse tree produced by FailureParser.
type FailureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFailure is called when entering the failure production.
	EnterFailure(c *FailureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFailure is called when exiting the failure production.
	ExitFailure(c *FailureContext)
}
