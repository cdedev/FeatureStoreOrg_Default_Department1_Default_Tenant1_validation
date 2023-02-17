// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676607095170/WeightGain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeightGain

import "github.com/antlr/antlr4/runtime/Go/antlr"

// WeightGainListener is a complete listener for a parse tree produced by WeightGainParser.
type WeightGainListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterWeightgain is called when entering the weightgain production.
	EnterWeightgain(c *WeightgainContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitWeightgain is called when exiting the weightgain production.
	ExitWeightgain(c *WeightgainContext)
}
