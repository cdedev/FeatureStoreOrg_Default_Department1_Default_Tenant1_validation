// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675069518379/Review.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Review

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ReviewListener is a complete listener for a parse tree produced by ReviewParser.
type ReviewListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterReview is called when entering the review production.
	EnterReview(c *ReviewContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitReview is called when exiting the review production.
	ExitReview(c *ReviewContext)
}
