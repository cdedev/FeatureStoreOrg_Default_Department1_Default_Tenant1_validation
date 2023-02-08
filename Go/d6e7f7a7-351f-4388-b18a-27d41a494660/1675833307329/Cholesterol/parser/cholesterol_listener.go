// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675833307329/Cholesterol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cholesterol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CholesterolListener is a complete listener for a parse tree produced by CholesterolParser.
type CholesterolListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCholesterol is called when entering the cholesterol production.
	EnterCholesterol(c *CholesterolContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCholesterol is called when exiting the cholesterol production.
	ExitCholesterol(c *CholesterolContext)
}
