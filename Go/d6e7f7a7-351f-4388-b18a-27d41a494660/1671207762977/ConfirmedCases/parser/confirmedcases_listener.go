// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671207762977/ConfirmedCases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ConfirmedCases

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConfirmedCasesListener is a complete listener for a parse tree produced by ConfirmedCasesParser.
type ConfirmedCasesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConfirmedcases is called when entering the confirmedcases production.
	EnterConfirmedcases(c *ConfirmedcasesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConfirmedcases is called when exiting the confirmedcases production.
	ExitConfirmedcases(c *ConfirmedcasesContext)
}
