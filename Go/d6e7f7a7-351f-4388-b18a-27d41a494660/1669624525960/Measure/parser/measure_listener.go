// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624525960/Measure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Measure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MeasureListener is a complete listener for a parse tree produced by MeasureParser.
type MeasureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMeasure is called when entering the measure production.
	EnterMeasure(c *MeasureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMeasure is called when exiting the measure production.
	ExitMeasure(c *MeasureContext)
}
