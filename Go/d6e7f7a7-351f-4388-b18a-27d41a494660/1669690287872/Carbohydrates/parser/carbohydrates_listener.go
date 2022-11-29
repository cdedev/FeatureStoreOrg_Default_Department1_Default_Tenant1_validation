// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669690287872/Carbohydrates.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carbohydrates

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CarbohydratesListener is a complete listener for a parse tree produced by CarbohydratesParser.
type CarbohydratesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCarbohydrates is called when entering the carbohydrates production.
	EnterCarbohydrates(c *CarbohydratesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCarbohydrates is called when exiting the carbohydrates production.
	ExitCarbohydrates(c *CarbohydratesContext)
}
