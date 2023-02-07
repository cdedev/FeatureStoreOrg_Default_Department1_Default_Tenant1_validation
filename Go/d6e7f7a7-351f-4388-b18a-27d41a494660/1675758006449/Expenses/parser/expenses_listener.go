// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675758006449/Expenses.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expenses

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExpensesListener is a complete listener for a parse tree produced by ExpensesParser.
type ExpensesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpenses is called when entering the expenses production.
	EnterExpenses(c *ExpensesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpenses is called when exiting the expenses production.
	ExitExpenses(c *ExpensesContext)
}
