// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669779000828/MultipleLines.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MultipleLines

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MultipleLinesListener is a complete listener for a parse tree produced by MultipleLinesParser.
type MultipleLinesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMultiplelines is called when entering the multiplelines production.
	EnterMultiplelines(c *MultiplelinesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMultiplelines is called when exiting the multiplelines production.
	ExitMultiplelines(c *MultiplelinesContext)
}
