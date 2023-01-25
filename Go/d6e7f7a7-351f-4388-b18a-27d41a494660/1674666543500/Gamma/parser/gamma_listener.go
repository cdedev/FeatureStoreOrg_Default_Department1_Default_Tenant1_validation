// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674666543500/Gamma.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gamma

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GammaListener is a complete listener for a parse tree produced by GammaParser.
type GammaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGamma is called when entering the gamma production.
	EnterGamma(c *GammaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGamma is called when exiting the gamma production.
	ExitGamma(c *GammaContext)
}
