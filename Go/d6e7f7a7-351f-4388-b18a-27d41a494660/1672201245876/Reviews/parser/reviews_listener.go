// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672201245876/Reviews.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Reviews

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReviewsListener is a complete listener for a parse tree produced by ReviewsParser.
type ReviewsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterReviews is called when entering the reviews production.
	EnterReviews(c *ReviewsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitReviews is called when exiting the reviews production.
	ExitReviews(c *ReviewsContext)
}
