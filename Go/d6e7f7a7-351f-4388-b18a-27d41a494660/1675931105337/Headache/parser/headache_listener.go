// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675931105337/Headache.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Headache

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HeadacheListener is a complete listener for a parse tree produced by HeadacheParser.
type HeadacheListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHeadache is called when entering the headache production.
	EnterHeadache(c *HeadacheContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHeadache is called when exiting the headache production.
	ExitHeadache(c *HeadacheContext)
}
