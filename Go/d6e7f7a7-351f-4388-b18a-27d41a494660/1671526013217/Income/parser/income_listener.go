// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671526013217/Income.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Income

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IncomeListener is a complete listener for a parse tree produced by IncomeParser.
type IncomeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIncome is called when entering the income production.
	EnterIncome(c *IncomeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIncome is called when exiting the income production.
	ExitIncome(c *IncomeContext)
}
