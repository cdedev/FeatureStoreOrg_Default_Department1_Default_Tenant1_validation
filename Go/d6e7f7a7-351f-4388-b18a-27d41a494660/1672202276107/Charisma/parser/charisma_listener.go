// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202276107/Charisma.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Charisma

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CharismaListener is a complete listener for a parse tree produced by CharismaParser.
type CharismaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCharisma is called when entering the charisma production.
	EnterCharisma(c *CharismaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCharisma is called when exiting the charisma production.
	ExitCharisma(c *CharismaContext)
}
