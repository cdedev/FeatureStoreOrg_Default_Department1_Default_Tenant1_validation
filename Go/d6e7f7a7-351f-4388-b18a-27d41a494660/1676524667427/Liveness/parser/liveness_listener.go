// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524667427/Liveness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Liveness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LivenessListener is a complete listener for a parse tree produced by LivenessParser.
type LivenessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLiveness is called when entering the liveness production.
	EnterLiveness(c *LivenessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLiveness is called when exiting the liveness production.
	ExitLiveness(c *LivenessContext)
}
