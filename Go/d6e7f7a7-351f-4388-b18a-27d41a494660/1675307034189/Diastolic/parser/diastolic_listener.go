// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675307034189/Diastolic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diastolic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiastolicListener is a complete listener for a parse tree produced by DiastolicParser.
type DiastolicListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDiastolic is called when entering the diastolic production.
	EnterDiastolic(c *DiastolicContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDiastolic is called when exiting the diastolic production.
	ExitDiastolic(c *DiastolicContext)
}
