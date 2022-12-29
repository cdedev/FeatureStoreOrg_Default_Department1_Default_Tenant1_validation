// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672287033447/Contact.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Contact

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ContactListener is a complete listener for a parse tree produced by ContactParser.
type ContactListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterContact is called when entering the contact production.
	EnterContact(c *ContactContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitContact is called when exiting the contact production.
	ExitContact(c *ContactContext)
}
