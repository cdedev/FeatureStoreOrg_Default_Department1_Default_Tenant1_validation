// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697774224/Alcohol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Alcohol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AlcoholListener is a complete listener for a parse tree produced by AlcoholParser.
type AlcoholListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAlcohol is called when entering the alcohol production.
	EnterAlcohol(c *AlcoholContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAlcohol is called when exiting the alcohol production.
	ExitAlcohol(c *AlcoholContext)
}
