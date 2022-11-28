// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653609323/Balance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Balance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BalanceListener is a complete listener for a parse tree produced by BalanceParser.
type BalanceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBalance is called when entering the balance production.
	EnterBalance(c *BalanceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBalance is called when exiting the balance production.
	ExitBalance(c *BalanceContext)
}
