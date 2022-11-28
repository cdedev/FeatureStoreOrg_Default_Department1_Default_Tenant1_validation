// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656096830/Respiration_rate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Respiration_rate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Respiration_rateListener is a complete listener for a parse tree produced by Respiration_rateParser.
type Respiration_rateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRespiration_rate is called when entering the respiration_rate production.
	EnterRespiration_rate(c *Respiration_rateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRespiration_rate is called when exiting the respiration_rate production.
	ExitRespiration_rate(c *Respiration_rateContext)
}
