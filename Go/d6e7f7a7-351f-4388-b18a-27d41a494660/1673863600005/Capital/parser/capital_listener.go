// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673863600005/Capital.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Capital

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CapitalListener is a complete listener for a parse tree produced by CapitalParser.
type CapitalListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCapital is called when entering the capital production.
	EnterCapital(c *CapitalContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCapital is called when exiting the capital production.
	ExitCapital(c *CapitalContext)
}
