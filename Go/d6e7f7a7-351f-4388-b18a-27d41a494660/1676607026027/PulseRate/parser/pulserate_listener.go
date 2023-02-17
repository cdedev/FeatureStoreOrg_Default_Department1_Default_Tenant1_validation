// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676607026027/PulseRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PulseRate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PulseRateListener is a complete listener for a parse tree produced by PulseRateParser.
type PulseRateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPulserate is called when entering the pulserate production.
	EnterPulserate(c *PulserateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPulserate is called when exiting the pulserate production.
	ExitPulserate(c *PulserateContext)
}
