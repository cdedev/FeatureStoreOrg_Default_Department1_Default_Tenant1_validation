// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673878334211/Phonophobia.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Phonophobia

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PhonophobiaListener is a complete listener for a parse tree produced by PhonophobiaParser.
type PhonophobiaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPhonophobia is called when entering the phonophobia production.
	EnterPhonophobia(c *PhonophobiaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPhonophobia is called when exiting the phonophobia production.
	ExitPhonophobia(c *PhonophobiaContext)
}
