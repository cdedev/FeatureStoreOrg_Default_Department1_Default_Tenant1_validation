// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235220015/Agreeableness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Agreeableness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AgreeablenessListener is a complete listener for a parse tree produced by AgreeablenessParser.
type AgreeablenessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAgreeableness is called when entering the agreeableness production.
	EnterAgreeableness(c *AgreeablenessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAgreeableness is called when exiting the agreeableness production.
	ExitAgreeableness(c *AgreeablenessContext)
}
