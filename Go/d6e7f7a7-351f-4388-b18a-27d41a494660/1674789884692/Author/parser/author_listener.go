// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674789884692/Author.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Author

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AuthorListener is a complete listener for a parse tree produced by AuthorParser.
type AuthorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAuthor is called when entering the author production.
	EnterAuthor(c *AuthorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAuthor is called when exiting the author production.
	ExitAuthor(c *AuthorContext)
}
