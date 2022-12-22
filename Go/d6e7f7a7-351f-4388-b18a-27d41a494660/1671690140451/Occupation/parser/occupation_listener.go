// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690140451/Occupation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Occupation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OccupationListener is a complete listener for a parse tree produced by OccupationParser.
type OccupationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOccupation is called when entering the occupation production.
	EnterOccupation(c *OccupationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOccupation is called when exiting the occupation production.
	ExitOccupation(c *OccupationContext)
}
