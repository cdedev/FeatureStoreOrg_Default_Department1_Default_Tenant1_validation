// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836976302/Period.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Period

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PeriodListener is a complete listener for a parse tree produced by PeriodParser.
type PeriodListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPeriod is called when entering the period production.
	EnterPeriod(c *PeriodContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPeriod is called when exiting the period production.
	ExitPeriod(c *PeriodContext)
}
