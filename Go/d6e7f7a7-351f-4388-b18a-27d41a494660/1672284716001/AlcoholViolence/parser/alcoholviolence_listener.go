// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284716001/AlcoholViolence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AlcoholViolence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AlcoholViolenceListener is a complete listener for a parse tree produced by AlcoholViolenceParser.
type AlcoholViolenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAlcoholviolence is called when entering the alcoholviolence production.
	EnterAlcoholviolence(c *AlcoholviolenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAlcoholviolence is called when exiting the alcoholviolence production.
	ExitAlcoholviolence(c *AlcoholviolenceContext)
}
