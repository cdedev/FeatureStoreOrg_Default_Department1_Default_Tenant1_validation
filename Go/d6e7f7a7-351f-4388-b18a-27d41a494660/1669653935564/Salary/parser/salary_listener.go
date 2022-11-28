// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653935564/Salary.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Salary

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SalaryListener is a complete listener for a parse tree produced by SalaryParser.
type SalaryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSalary is called when entering the salary production.
	EnterSalary(c *SalaryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSalary is called when exiting the salary production.
	ExitSalary(c *SalaryContext)
}
