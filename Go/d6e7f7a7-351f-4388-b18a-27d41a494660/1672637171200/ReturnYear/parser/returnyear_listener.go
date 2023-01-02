// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672637171200/ReturnYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ReturnYear

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReturnYearListener is a complete listener for a parse tree produced by ReturnYearParser.
type ReturnYearListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterReturnyear is called when entering the returnyear production.
	EnterReturnyear(c *ReturnyearContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitReturnyear is called when exiting the returnyear production.
	ExitReturnyear(c *ReturnyearContext)
}
