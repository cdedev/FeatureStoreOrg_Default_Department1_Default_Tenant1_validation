// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671691110785/Transaction.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Transaction

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TransactionListener is a complete listener for a parse tree produced by TransactionParser.
type TransactionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTransaction is called when entering the transaction production.
	EnterTransaction(c *TransactionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTransaction is called when exiting the transaction production.
	ExitTransaction(c *TransactionContext)
}
