// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674541115725/Married.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Married

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MarriedListener is a complete listener for a parse tree produced by MarriedParser.
type MarriedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMarried is called when entering the married production.
	EnterMarried(c *MarriedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMarried is called when exiting the married production.
	ExitMarried(c *MarriedContext)
}
