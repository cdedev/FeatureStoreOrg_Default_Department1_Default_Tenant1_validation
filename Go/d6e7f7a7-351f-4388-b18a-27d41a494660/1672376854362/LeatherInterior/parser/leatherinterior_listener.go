// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376854362/LeatherInterior.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LeatherInterior

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LeatherInteriorListener is a complete listener for a parse tree produced by LeatherInteriorParser.
type LeatherInteriorListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLeatherinterior is called when entering the leatherinterior production.
	EnterLeatherinterior(c *LeatherinteriorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLeatherinterior is called when exiting the leatherinterior production.
	ExitLeatherinterior(c *LeatherinteriorContext)
}
