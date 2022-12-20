// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671528831254/LoanTerm.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LoanTerm

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LoanTermListener is a complete listener for a parse tree produced by LoanTermParser.
type LoanTermListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLoanterm is called when entering the loanterm production.
	EnterLoanterm(c *LoantermContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLoanterm is called when exiting the loanterm production.
	ExitLoanterm(c *LoantermContext)
}
