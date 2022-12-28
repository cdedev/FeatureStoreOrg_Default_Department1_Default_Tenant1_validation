// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672202956618/NewBalanceOrig.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NewBalanceOrig

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NewBalanceOrigListener is a complete listener for a parse tree produced by NewBalanceOrigParser.
type NewBalanceOrigListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNewbalanceorig is called when entering the newbalanceorig production.
	EnterNewbalanceorig(c *NewbalanceorigContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNewbalanceorig is called when exiting the newbalanceorig production.
	ExitNewbalanceorig(c *NewbalanceorigContext)
}
