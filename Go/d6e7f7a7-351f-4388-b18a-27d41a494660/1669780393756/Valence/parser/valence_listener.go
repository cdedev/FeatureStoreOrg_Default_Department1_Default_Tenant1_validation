// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780393756/Valence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Valence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ValenceListener is a complete listener for a parse tree produced by ValenceParser.
type ValenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterValence is called when entering the valence production.
	EnterValence(c *ValenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitValence is called when exiting the valence production.
	ExitValence(c *ValenceContext)
}
