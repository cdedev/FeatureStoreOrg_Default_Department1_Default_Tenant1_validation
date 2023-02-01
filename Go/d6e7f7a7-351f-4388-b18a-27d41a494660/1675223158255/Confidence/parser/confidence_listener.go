// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675223158255/Confidence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Confidence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConfidenceListener is a complete listener for a parse tree produced by ConfidenceParser.
type ConfidenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConfidence is called when entering the confidence production.
	EnterConfidence(c *ConfidenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConfidence is called when exiting the confidence production.
	ExitConfidence(c *ConfidenceContext)
}
