// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076407953/WaterSaturationFraction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WaterSaturationFraction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WaterSaturationFractionListener is a complete listener for a parse tree produced by WaterSaturationFractionParser.
type WaterSaturationFractionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWatersaturationfraction is called when entering the watersaturationfraction production.
	EnterWatersaturationfraction(c *WatersaturationfractionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWatersaturationfraction is called when exiting the watersaturationfraction production.
	ExitWatersaturationfraction(c *WatersaturationfractionContext)
}
