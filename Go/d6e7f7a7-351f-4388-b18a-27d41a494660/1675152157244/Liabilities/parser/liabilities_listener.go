// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675152157244/Liabilities.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Liabilities

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LiabilitiesListener is a complete listener for a parse tree produced by LiabilitiesParser.
type LiabilitiesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLiabilities is called when entering the liabilities production.
	EnterLiabilities(c *LiabilitiesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLiabilities is called when exiting the liabilities production.
	ExitLiabilities(c *LiabilitiesContext)
}
