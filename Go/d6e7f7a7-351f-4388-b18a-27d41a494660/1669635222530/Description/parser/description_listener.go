// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669635222530/Description.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Description

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DescriptionListener is a complete listener for a parse tree produced by DescriptionParser.
type DescriptionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDescription is called when entering the description production.
	EnterDescription(c *DescriptionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDescription is called when exiting the description production.
	ExitDescription(c *DescriptionContext)
}
