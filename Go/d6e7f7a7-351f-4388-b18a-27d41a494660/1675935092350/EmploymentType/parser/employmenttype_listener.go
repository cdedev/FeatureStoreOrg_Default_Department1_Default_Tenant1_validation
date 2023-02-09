// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675935092350/EmploymentType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EmploymentType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EmploymentTypeListener is a complete listener for a parse tree produced by EmploymentTypeParser.
type EmploymentTypeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEmploymenttype is called when entering the employmenttype production.
	EnterEmploymenttype(c *EmploymenttypeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEmploymenttype is called when exiting the employmenttype production.
	ExitEmploymenttype(c *EmploymenttypeContext)
}
