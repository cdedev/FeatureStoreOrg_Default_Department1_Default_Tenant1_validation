// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096810128/Pitch.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pitch

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PitchListener is a complete listener for a parse tree produced by PitchParser.
type PitchListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPitch is called when entering the pitch production.
	EnterPitch(c *PitchContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPitch is called when exiting the pitch production.
	ExitPitch(c *PitchContext)
}
