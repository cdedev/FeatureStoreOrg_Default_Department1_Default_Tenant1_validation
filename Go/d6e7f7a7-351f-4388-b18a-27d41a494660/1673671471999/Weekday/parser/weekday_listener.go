// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673671471999/Weekday.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Weekday

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WeekdayListener is a complete listener for a parse tree produced by WeekdayParser.
type WeekdayListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWeekday is called when entering the weekday production.
	EnterWeekday(c *WeekdayContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWeekday is called when exiting the weekday production.
	ExitWeekday(c *WeekdayContext)
}
