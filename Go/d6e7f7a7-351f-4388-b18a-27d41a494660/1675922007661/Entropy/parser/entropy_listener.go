// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675922007661/Entropy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Entropy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EntropyListener is a complete listener for a parse tree produced by EntropyParser.
type EntropyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEntropy is called when entering the entropy production.
	EnterEntropy(c *EntropyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEntropy is called when exiting the entropy production.
	ExitEntropy(c *EntropyContext)
}
