// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674014957141/Jaundice.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Jaundice

import "github.com/antlr/antlr4/runtime/Go/antlr"

// JaundiceListener is a complete listener for a parse tree produced by JaundiceParser.
type JaundiceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJaundice is called when entering the jaundice production.
	EnterJaundice(c *JaundiceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJaundice is called when exiting the jaundice production.
	ExitJaundice(c *JaundiceContext)
}
