// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672201967666/Utterance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Utterance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// UtteranceListener is a complete listener for a parse tree produced by UtteranceParser.
type UtteranceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUtterance is called when entering the utterance production.
	EnterUtterance(c *UtteranceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUtterance is called when exiting the utterance production.
	ExitUtterance(c *UtteranceContext)
}
