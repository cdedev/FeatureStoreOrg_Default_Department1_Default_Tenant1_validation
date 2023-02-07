// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675745067273/HeadQuarters.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HeadQuarters

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HeadQuartersListener is a complete listener for a parse tree produced by HeadQuartersParser.
type HeadQuartersListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHeadquarters is called when entering the headquarters production.
	EnterHeadquarters(c *HeadquartersContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHeadquarters is called when exiting the headquarters production.
	ExitHeadquarters(c *HeadquartersContext)
}
