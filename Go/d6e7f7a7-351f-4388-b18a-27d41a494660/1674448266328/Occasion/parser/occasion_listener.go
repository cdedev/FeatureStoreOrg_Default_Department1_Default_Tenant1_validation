// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674448266328/Occasion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Occasion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OccasionListener is a complete listener for a parse tree produced by OccasionParser.
type OccasionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOccasion is called when entering the occasion production.
	EnterOccasion(c *OccasionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOccasion is called when exiting the occasion production.
	ExitOccasion(c *OccasionContext)
}
