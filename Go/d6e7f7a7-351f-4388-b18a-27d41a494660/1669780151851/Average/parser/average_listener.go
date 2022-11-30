// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780151851/Average.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Average

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AverageListener is a complete listener for a parse tree produced by AverageParser.
type AverageListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAverage is called when entering the average production.
	EnterAverage(c *AverageContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAverage is called when exiting the average production.
	ExitAverage(c *AverageContext)
}
