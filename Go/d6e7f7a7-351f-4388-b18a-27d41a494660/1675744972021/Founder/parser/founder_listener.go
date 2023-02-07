// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675744972021/Founder.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Founder

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FounderListener is a complete listener for a parse tree produced by FounderParser.
type FounderListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFounder is called when entering the founder production.
	EnterFounder(c *FounderContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFounder is called when exiting the founder production.
	ExitFounder(c *FounderContext)
}
