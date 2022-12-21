// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671597462999/ResidenceRegion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ResidenceRegion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ResidenceRegionListener is a complete listener for a parse tree produced by ResidenceRegionParser.
type ResidenceRegionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterResidenceregion is called when entering the residenceregion production.
	EnterResidenceregion(c *ResidenceregionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitResidenceregion is called when exiting the residenceregion production.
	ExitResidenceregion(c *ResidenceregionContext)
}
