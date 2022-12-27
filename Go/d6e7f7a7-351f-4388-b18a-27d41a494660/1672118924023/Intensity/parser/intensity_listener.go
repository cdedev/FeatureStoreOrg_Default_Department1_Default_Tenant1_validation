// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118924023/Intensity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Intensity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IntensityListener is a complete listener for a parse tree produced by IntensityParser.
type IntensityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIntensity is called when entering the intensity production.
	EnterIntensity(c *IntensityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIntensity is called when exiting the intensity production.
	ExitIntensity(c *IntensityContext)
}
