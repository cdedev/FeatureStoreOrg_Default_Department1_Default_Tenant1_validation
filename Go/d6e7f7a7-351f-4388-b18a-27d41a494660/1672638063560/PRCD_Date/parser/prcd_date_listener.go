// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672638063560/PRCD_Date.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PRCD_Date

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PRCD_DateListener is a complete listener for a parse tree produced by PRCD_DateParser.
type PRCD_DateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPrcd_date is called when entering the prcd_date production.
	EnterPrcd_date(c *Prcd_dateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPrcd_date is called when exiting the prcd_date production.
	ExitPrcd_date(c *Prcd_dateContext)
}
