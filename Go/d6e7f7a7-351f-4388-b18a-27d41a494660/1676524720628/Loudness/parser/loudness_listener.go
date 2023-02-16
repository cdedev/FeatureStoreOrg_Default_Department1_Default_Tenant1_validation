// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524720628/Loudness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Loudness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LoudnessListener is a complete listener for a parse tree produced by LoudnessParser.
type LoudnessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLoudness is called when entering the loudness production.
	EnterLoudness(c *LoudnessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLoudness is called when exiting the loudness production.
	ExitLoudness(c *LoudnessContext)
}
