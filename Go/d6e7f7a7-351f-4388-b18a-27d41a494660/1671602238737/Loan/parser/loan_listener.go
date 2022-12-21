// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671602238737/Loan.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loan

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LoanListener is a complete listener for a parse tree produced by LoanParser.
type LoanListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLoan is called when entering the loan production.
	EnterLoan(c *LoanContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLoan is called when exiting the loan production.
	ExitLoan(c *LoanContext)
}
