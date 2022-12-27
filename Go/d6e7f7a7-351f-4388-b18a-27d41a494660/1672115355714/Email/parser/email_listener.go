// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115355714/Email.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Email

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EmailListener is a complete listener for a parse tree produced by EmailParser.
type EmailListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEmail is called when entering the email production.
	EnterEmail(c *EmailContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEmail is called when exiting the email production.
	ExitEmail(c *EmailContext)
}
