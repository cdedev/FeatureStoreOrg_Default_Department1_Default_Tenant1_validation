// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656174771/Sleeping_hours.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sleeping_hours

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Sleeping_hoursListener is a complete listener for a parse tree produced by Sleeping_hoursParser.
type Sleeping_hoursListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSleeping_hours is called when entering the sleeping_hours production.
	EnterSleeping_hours(c *Sleeping_hoursContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSleeping_hours is called when exiting the sleeping_hours production.
	ExitSleeping_hours(c *Sleeping_hoursContext)
}
