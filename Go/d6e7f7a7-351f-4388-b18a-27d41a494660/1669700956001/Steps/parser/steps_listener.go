// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700956001/Steps.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Steps

import "github.com/antlr/antlr4/runtime/Go/antlr"

// StepsListener is a complete listener for a parse tree produced by StepsParser.
type StepsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSteps is called when entering the steps production.
	EnterSteps(c *StepsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSteps is called when exiting the steps production.
	ExitSteps(c *StepsContext)
}
