// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672114544968/Interval.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Interval

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IntervalListener is a complete listener for a parse tree produced by IntervalParser.
type IntervalListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)
}
