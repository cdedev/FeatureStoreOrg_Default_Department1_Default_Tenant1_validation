// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674538873750/Noise.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Noise

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NoiseListener is a complete listener for a parse tree produced by NoiseParser.
type NoiseListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNoise is called when entering the noise production.
	EnterNoise(c *NoiseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNoise is called when exiting the noise production.
	ExitNoise(c *NoiseContext)
}
