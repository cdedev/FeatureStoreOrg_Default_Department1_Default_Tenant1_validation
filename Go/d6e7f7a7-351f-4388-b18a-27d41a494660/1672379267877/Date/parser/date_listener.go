// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672379267877/Date.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Date

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DateListener is a complete listener for a parse tree produced by DateParser.
type DateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)
}
