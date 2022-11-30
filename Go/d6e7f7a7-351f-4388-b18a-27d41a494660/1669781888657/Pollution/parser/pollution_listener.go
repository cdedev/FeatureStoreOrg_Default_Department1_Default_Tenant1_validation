// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669781888657/Pollution.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pollution

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PollutionListener is a complete listener for a parse tree produced by PollutionParser.
type PollutionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPollution is called when entering the pollution production.
	EnterPollution(c *PollutionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPollution is called when exiting the pollution production.
	ExitPollution(c *PollutionContext)
}
