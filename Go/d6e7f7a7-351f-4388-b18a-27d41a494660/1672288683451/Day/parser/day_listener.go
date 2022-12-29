// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288683451/Day.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Day

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DayListener is a complete listener for a parse tree produced by DayParser.
type DayListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDay is called when entering the day production.
	EnterDay(c *DayContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDay is called when exiting the day production.
	ExitDay(c *DayContext)
}
