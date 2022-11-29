// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697912745/Density.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Density

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DensityListener is a complete listener for a parse tree produced by DensityParser.
type DensityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDensity is called when entering the density production.
	EnterDensity(c *DensityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDensity is called when exiting the density production.
	ExitDensity(c *DensityContext)
}
