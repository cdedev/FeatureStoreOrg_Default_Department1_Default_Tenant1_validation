// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673928568160/Postcode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Postcode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PostcodeListener is a complete listener for a parse tree produced by PostcodeParser.
type PostcodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPostcode is called when entering the postcode production.
	EnterPostcode(c *PostcodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPostcode is called when exiting the postcode production.
	ExitPostcode(c *PostcodeContext)
}
