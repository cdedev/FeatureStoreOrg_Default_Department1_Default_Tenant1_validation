// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671193515509/ReviewsCount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ReviewsCount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReviewsCountListener is a complete listener for a parse tree produced by ReviewsCountParser.
type ReviewsCountListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterReviewscount is called when entering the reviewscount production.
	EnterReviewscount(c *ReviewscountContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitReviewscount is called when exiting the reviewscount production.
	ExitReviewscount(c *ReviewscountContext)
}
