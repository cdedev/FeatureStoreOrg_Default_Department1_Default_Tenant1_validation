// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672810323495/Circumference.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Circumference

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CircumferenceListener is a complete listener for a parse tree produced by CircumferenceParser.
type CircumferenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCircumference is called when entering the circumference production.
	EnterCircumference(c *CircumferenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCircumference is called when exiting the circumference production.
	ExitCircumference(c *CircumferenceContext)
}
