// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672638199157/Trans_Date.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Trans_Date

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Trans_DateListener is a complete listener for a parse tree produced by Trans_DateParser.
type Trans_DateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTrans_date is called when entering the trans_date production.
	EnterTrans_date(c *Trans_dateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTrans_date is called when exiting the trans_date production.
	ExitTrans_date(c *Trans_dateContext)
}
