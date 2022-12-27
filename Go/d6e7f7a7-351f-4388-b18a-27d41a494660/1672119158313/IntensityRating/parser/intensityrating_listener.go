// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672119158313/IntensityRating.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntensityRating

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IntensityRatingListener is a complete listener for a parse tree produced by IntensityRatingParser.
type IntensityRatingListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIntensityrating is called when entering the intensityrating production.
	EnterIntensityrating(c *IntensityratingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIntensityrating is called when exiting the intensityrating production.
	ExitIntensityrating(c *IntensityratingContext)
}
