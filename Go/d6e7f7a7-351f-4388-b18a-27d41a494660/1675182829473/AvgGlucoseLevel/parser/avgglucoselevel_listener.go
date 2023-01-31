// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675182829473/AvgGlucoseLevel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AvgGlucoseLevel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AvgGlucoseLevelListener is a complete listener for a parse tree produced by AvgGlucoseLevelParser.
type AvgGlucoseLevelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAvgglucoselevel is called when entering the avgglucoselevel production.
	EnterAvgglucoselevel(c *AvgglucoselevelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAvgglucoselevel is called when exiting the avgglucoselevel production.
	ExitAvgglucoselevel(c *AvgglucoselevelContext)
}
