// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624790720/City.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // City

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CityListener is a complete listener for a parse tree produced by CityParser.
type CityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCity is called when entering the city production.
	EnterCity(c *CityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCity is called when exiting the city production.
	ExitCity(c *CityContext)
}
