// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675934751714/Insurance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Insurance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// InsuranceListener is a complete listener for a parse tree produced by InsuranceParser.
type InsuranceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterInsurance is called when entering the insurance production.
	EnterInsurance(c *InsuranceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitInsurance is called when exiting the insurance production.
	ExitInsurance(c *InsuranceContext)
}
