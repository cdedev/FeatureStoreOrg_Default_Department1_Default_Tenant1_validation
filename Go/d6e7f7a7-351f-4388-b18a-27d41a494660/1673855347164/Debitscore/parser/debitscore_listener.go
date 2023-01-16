// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673855347164/Debitscore.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Debitscore

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DebitscoreListener is a complete listener for a parse tree produced by DebitscoreParser.
type DebitscoreListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDebitscore is called when entering the debitscore production.
	EnterDebitscore(c *DebitscoreContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDebitscore is called when exiting the debitscore production.
	ExitDebitscore(c *DebitscoreContext)
}
