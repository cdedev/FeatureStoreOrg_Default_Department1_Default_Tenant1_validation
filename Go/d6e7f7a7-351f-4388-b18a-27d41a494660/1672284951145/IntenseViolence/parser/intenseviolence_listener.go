// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284951145/IntenseViolence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntenseViolence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IntenseViolenceListener is a complete listener for a parse tree produced by IntenseViolenceParser.
type IntenseViolenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIntenseviolence is called when entering the intenseviolence production.
	EnterIntenseviolence(c *IntenseviolenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIntenseviolence is called when exiting the intenseviolence production.
	ExitIntenseviolence(c *IntenseviolenceContext)
}
