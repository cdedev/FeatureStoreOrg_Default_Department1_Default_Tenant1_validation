// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524537633/Acousticness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Acousticness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AcousticnessListener is a complete listener for a parse tree produced by AcousticnessParser.
type AcousticnessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAcousticness is called when entering the acousticness production.
	EnterAcousticness(c *AcousticnessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAcousticness is called when exiting the acousticness production.
	ExitAcousticness(c *AcousticnessContext)
}
