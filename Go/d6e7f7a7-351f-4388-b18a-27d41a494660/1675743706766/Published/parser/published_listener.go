// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675743706766/Published.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Published

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PublishedListener is a complete listener for a parse tree produced by PublishedParser.
type PublishedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPublished is called when entering the published production.
	EnterPublished(c *PublishedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPublished is called when exiting the published production.
	ExitPublished(c *PublishedContext)
}
