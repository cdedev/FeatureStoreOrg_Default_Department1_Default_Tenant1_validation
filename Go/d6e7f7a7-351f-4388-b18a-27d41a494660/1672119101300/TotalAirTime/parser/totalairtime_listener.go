// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119101300/TotalAirTime.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TotalAirTime

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TotalAirTimeListener is a complete listener for a parse tree produced by TotalAirTimeParser.
type TotalAirTimeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTotalairtime is called when entering the totalairtime production.
	EnterTotalairtime(c *TotalairtimeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTotalairtime is called when exiting the totalairtime production.
	ExitTotalairtime(c *TotalairtimeContext)
}
