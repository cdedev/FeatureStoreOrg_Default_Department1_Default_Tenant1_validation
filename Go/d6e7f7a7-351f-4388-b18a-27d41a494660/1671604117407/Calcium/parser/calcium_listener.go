// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604117407/Calcium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Calcium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalciumListener is a complete listener for a parse tree produced by CalciumParser.
type CalciumListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCalcium is called when entering the calcium production.
	EnterCalcium(c *CalciumContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCalcium is called when exiting the calcium production.
	ExitCalcium(c *CalciumContext)
}
