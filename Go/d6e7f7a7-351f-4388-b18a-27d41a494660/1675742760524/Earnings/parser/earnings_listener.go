// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675742760524/Earnings.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Earnings

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EarningsListener is a complete listener for a parse tree produced by EarningsParser.
type EarningsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEarnings is called when entering the earnings production.
	EnterEarnings(c *EarningsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEarnings is called when exiting the earnings production.
	ExitEarnings(c *EarningsContext)
}
