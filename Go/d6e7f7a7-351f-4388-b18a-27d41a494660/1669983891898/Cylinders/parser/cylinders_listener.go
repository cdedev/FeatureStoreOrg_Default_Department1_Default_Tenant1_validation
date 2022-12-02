// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669983891898/Cylinders.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cylinders

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CylindersListener is a complete listener for a parse tree produced by CylindersParser.
type CylindersListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCylinders is called when entering the cylinders production.
	EnterCylinders(c *CylindersContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCylinders is called when exiting the cylinders production.
	ExitCylinders(c *CylindersContext)
}
