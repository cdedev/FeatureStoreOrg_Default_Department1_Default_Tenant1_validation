// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673083863888/Subregion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Subregion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SubregionListener is a complete listener for a parse tree produced by SubregionParser.
type SubregionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSubregion is called when entering the subregion production.
	EnterSubregion(c *SubregionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSubregion is called when exiting the subregion production.
	ExitSubregion(c *SubregionContext)
}
