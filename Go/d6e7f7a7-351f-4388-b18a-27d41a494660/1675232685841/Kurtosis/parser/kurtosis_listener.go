// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675232685841/Kurtosis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Kurtosis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// KurtosisListener is a complete listener for a parse tree produced by KurtosisParser.
type KurtosisListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterKurtosis is called when entering the kurtosis production.
	EnterKurtosis(c *KurtosisContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitKurtosis is called when exiting the kurtosis production.
	ExitKurtosis(c *KurtosisContext)
}
