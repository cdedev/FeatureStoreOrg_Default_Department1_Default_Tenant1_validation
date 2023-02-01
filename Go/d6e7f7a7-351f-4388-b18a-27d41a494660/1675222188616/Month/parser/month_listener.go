// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675222188616/Month.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Month

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MonthListener is a complete listener for a parse tree produced by MonthParser.
type MonthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMonth is called when entering the month production.
	EnterMonth(c *MonthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMonth is called when exiting the month production.
	ExitMonth(c *MonthContext)
}
