// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673242358580/State.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // State

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StateListener is a complete listener for a parse tree produced by StateParser.
type StateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterState is called when entering the state production.
	EnterState(c *StateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitState is called when exiting the state production.
	ExitState(c *StateContext)
}
