// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674837144984/School.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // School

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SchoolListener is a complete listener for a parse tree produced by SchoolParser.
type SchoolListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSchool is called when entering the school production.
	EnterSchool(c *SchoolContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSchool is called when exiting the school production.
	ExitSchool(c *SchoolContext)
}
