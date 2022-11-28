// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624297226/Frequency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Frequency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FrequencyListener is a complete listener for a parse tree produced by FrequencyParser.
type FrequencyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFrequency is called when entering the frequency production.
	EnterFrequency(c *FrequencyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFrequency is called when exiting the frequency production.
	ExitFrequency(c *FrequencyContext)
}
