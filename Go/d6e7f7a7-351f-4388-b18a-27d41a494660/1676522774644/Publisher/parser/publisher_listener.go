// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522774644/Publisher.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Publisher

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PublisherListener is a complete listener for a parse tree produced by PublisherParser.
type PublisherListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPublisher is called when entering the publisher production.
	EnterPublisher(c *PublisherContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPublisher is called when exiting the publisher production.
	ExitPublisher(c *PublisherContext)
}
