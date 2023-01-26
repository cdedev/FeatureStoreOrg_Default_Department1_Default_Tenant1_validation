// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674718945672/Employer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Employer

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EmployerListener is a complete listener for a parse tree produced by EmployerParser.
type EmployerListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEmployer is called when entering the employer production.
	EnterEmployer(c *EmployerContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEmployer is called when exiting the employer production.
	ExitEmployer(c *EmployerContext)
}
