// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675307081052/Systolic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Systolic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SystolicListener is a complete listener for a parse tree produced by SystolicParser.
type SystolicListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSystolic is called when entering the systolic production.
	EnterSystolic(c *SystolicContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSystolic is called when exiting the systolic production.
	ExitSystolic(c *SystolicContext)
}
