// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524772055/Popularity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Popularity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PopularityListener is a complete listener for a parse tree produced by PopularityParser.
type PopularityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPopularity is called when entering the popularity production.
	EnterPopularity(c *PopularityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPopularity is called when exiting the popularity production.
	ExitPopularity(c *PopularityContext)
}
