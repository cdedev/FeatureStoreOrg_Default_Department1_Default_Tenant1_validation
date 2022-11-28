// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624662791/Time.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Time

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TimeListener is a complete listener for a parse tree produced by TimeParser.
type TimeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTime is called when entering the time production.
	EnterTime(c *TimeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTime is called when exiting the time production.
	ExitTime(c *TimeContext)
}
