// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672977218378/WeekdayName.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeekdayName

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WeekdayNameListener is a complete listener for a parse tree produced by WeekdayNameParser.
type WeekdayNameListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWeekdayname is called when entering the weekdayname production.
	EnterWeekdayname(c *WeekdaynameContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWeekdayname is called when exiting the weekdayname production.
	ExitWeekdayname(c *WeekdaynameContext)
}
