// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674667997673/Debt.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Debt

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DebtListener is a complete listener for a parse tree produced by DebtParser.
type DebtListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDebt is called when entering the debt production.
	EnterDebt(c *DebtContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDebt is called when exiting the debt production.
	ExitDebt(c *DebtContext)
}
