// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884881303/AvgNoiseLvl.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AvgNoiseLvl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AvgNoiseLvlListener is a complete listener for a parse tree produced by AvgNoiseLvlParser.
type AvgNoiseLvlListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAvgnoiselvl is called when entering the avgnoiselvl production.
	EnterAvgnoiselvl(c *AvgnoiselvlContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAvgnoiselvl is called when exiting the avgnoiselvl production.
	ExitAvgnoiselvl(c *AvgnoiselvlContext)
}
