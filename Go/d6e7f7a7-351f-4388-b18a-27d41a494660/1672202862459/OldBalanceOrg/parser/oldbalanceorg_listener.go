// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202862459/OldBalanceOrg.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OldBalanceOrg

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OldBalanceOrgListener is a complete listener for a parse tree produced by OldBalanceOrgParser.
type OldBalanceOrgListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOldbalanceorg is called when entering the oldbalanceorg production.
	EnterOldbalanceorg(c *OldbalanceorgContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOldbalanceorg is called when exiting the oldbalanceorg production.
	ExitOldbalanceorg(c *OldbalanceorgContext)
}
