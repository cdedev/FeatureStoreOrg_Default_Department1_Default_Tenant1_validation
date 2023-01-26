// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719079626/Housing.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Housing

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HousingListener is a complete listener for a parse tree produced by HousingParser.
type HousingListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHousing is called when entering the housing production.
	EnterHousing(c *HousingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHousing is called when exiting the housing production.
	ExitHousing(c *HousingContext)
}
