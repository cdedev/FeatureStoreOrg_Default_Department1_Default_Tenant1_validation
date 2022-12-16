// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671205043399/Birth_Place.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Birth_Place

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Birth_PlaceListener is a complete listener for a parse tree produced by Birth_PlaceParser.
type Birth_PlaceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBirth_place is called when entering the birth_place production.
	EnterBirth_place(c *Birth_placeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBirth_place is called when exiting the birth_place production.
	ExitBirth_place(c *Birth_placeContext)
}
