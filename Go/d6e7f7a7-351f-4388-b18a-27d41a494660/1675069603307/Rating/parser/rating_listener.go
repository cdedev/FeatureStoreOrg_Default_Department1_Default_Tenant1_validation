// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675069603307/Rating.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rating

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RatingListener is a complete listener for a parse tree produced by RatingParser.
type RatingListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRating is called when entering the rating production.
	EnterRating(c *RatingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRating is called when exiting the rating production.
	ExitRating(c *RatingContext)
}
