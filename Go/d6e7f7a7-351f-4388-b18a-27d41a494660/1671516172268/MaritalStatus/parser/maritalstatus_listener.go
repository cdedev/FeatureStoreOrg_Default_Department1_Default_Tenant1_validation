// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671516172268/MaritalStatus.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaritalStatus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MaritalStatusListener is a complete listener for a parse tree produced by MaritalStatusParser.
type MaritalStatusListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMaritalstatus is called when entering the maritalstatus production.
	EnterMaritalstatus(c *MaritalstatusContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMaritalstatus is called when exiting the maritalstatus production.
	ExitMaritalstatus(c *MaritalstatusContext)
}
