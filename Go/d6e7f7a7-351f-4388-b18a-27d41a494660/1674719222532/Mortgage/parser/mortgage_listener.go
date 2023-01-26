// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719222532/Mortgage.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Mortgage

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MortgageListener is a complete listener for a parse tree produced by MortgageParser.
type MortgageListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMortgage is called when entering the mortgage production.
	EnterMortgage(c *MortgageContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMortgage is called when exiting the mortgage production.
	ExitMortgage(c *MortgageContext)
}
