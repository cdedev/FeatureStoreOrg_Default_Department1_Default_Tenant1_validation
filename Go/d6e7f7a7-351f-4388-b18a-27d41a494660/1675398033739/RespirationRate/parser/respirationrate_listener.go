// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675398033739/RespirationRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RespirationRate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RespirationRateListener is a complete listener for a parse tree produced by RespirationRateParser.
type RespirationRateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRespirationrate is called when entering the respirationrate production.
	EnterRespirationrate(c *RespirationrateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRespirationrate is called when exiting the respirationrate production.
	ExitRespirationrate(c *RespirationrateContext)
}
