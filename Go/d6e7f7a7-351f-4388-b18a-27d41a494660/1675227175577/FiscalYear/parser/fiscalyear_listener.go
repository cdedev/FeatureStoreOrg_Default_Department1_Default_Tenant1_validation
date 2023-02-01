// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675227175577/FiscalYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FiscalYear

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FiscalYearListener is a complete listener for a parse tree produced by FiscalYearParser.
type FiscalYearListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFiscalyear is called when entering the fiscalyear production.
	EnterFiscalyear(c *FiscalyearContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFiscalyear is called when exiting the fiscalyear production.
	ExitFiscalyear(c *FiscalyearContext)
}
