// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672805125776/Hypothesis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hypothesis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HypothesisListener is a complete listener for a parse tree produced by HypothesisParser.
type HypothesisListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHypothesis is called when entering the hypothesis production.
	EnterHypothesis(c *HypothesisContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHypothesis is called when exiting the hypothesis production.
	ExitHypothesis(c *HypothesisContext)
}
