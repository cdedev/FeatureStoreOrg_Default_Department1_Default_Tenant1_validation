// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671698471984/TotalWorkingYears.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TotalWorkingYears

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TotalWorkingYearsListener is a complete listener for a parse tree produced by TotalWorkingYearsParser.
type TotalWorkingYearsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTotalworkingyears is called when entering the totalworkingyears production.
	EnterTotalworkingyears(c *TotalworkingyearsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTotalworkingyears is called when exiting the totalworkingyears production.
	ExitTotalworkingyears(c *TotalworkingyearsContext)
}
