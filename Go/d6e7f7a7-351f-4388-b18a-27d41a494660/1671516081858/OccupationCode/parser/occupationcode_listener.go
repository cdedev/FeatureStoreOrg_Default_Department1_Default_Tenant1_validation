// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671516081858/OccupationCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OccupationCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OccupationCodeListener is a complete listener for a parse tree produced by OccupationCodeParser.
type OccupationCodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOccupationcode is called when entering the occupationcode production.
	EnterOccupationcode(c *OccupationcodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOccupationcode is called when exiting the occupationcode production.
	ExitOccupationcode(c *OccupationcodeContext)
}
