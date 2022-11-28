// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653702314/Department.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Department

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DepartmentListener is a complete listener for a parse tree produced by DepartmentParser.
type DepartmentListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDepartment is called when entering the department production.
	EnterDepartment(c *DepartmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDepartment is called when exiting the department production.
	ExitDepartment(c *DepartmentContext)
}
