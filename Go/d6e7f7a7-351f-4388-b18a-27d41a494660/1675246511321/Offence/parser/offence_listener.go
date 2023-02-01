// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675246511321/Offence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Offence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OffenceListener is a complete listener for a parse tree produced by OffenceParser.
type OffenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOffence is called when entering the offence production.
	EnterOffence(c *OffenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOffence is called when exiting the offence production.
	ExitOffence(c *OffenceContext)
}
