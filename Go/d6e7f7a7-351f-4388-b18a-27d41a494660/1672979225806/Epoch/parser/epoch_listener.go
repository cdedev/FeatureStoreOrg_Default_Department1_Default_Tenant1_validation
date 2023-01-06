// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979225806/Epoch.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Epoch

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EpochListener is a complete listener for a parse tree produced by EpochParser.
type EpochListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEpoch is called when entering the epoch production.
	EnterEpoch(c *EpochContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEpoch is called when exiting the epoch production.
	ExitEpoch(c *EpochContext)
}
