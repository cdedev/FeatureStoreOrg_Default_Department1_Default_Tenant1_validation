// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447352788/Baseline.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Baseline

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaselineListener is a complete listener for a parse tree produced by BaselineParser.
type BaselineListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBaseline is called when entering the baseline production.
	EnterBaseline(c *BaselineContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBaseline is called when exiting the baseline production.
	ExitBaseline(c *BaselineContext)
}
