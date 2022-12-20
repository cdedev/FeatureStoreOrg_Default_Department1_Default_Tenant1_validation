// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671514143700/PostalCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PostalCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PostalCodeListener is a complete listener for a parse tree produced by PostalCodeParser.
type PostalCodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPostalcode is called when entering the postalcode production.
	EnterPostalcode(c *PostalcodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPostalcode is called when exiting the postalcode production.
	ExitPostalcode(c *PostalcodeContext)
}
