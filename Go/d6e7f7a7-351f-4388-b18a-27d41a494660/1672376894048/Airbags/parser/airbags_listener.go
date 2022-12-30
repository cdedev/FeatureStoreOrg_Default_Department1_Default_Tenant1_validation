// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376894048/Airbags.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Airbags

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AirbagsListener is a complete listener for a parse tree produced by AirbagsParser.
type AirbagsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAirbags is called when entering the airbags production.
	EnterAirbags(c *AirbagsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAirbags is called when exiting the airbags production.
	ExitAirbags(c *AirbagsContext)
}
