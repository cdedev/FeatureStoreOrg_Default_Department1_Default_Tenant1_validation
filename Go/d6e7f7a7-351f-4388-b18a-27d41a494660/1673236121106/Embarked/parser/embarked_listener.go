// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673236121106/Embarked.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Embarked

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EmbarkedListener is a complete listener for a parse tree produced by EmbarkedParser.
type EmbarkedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEmbarked is called when entering the embarked production.
	EnterEmbarked(c *EmbarkedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEmbarked is called when exiting the embarked production.
	ExitEmbarked(c *EmbarkedContext)
}
