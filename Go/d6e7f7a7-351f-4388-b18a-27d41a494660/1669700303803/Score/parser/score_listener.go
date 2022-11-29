// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700303803/Score.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Score

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ScoreListener is a complete listener for a parse tree produced by ScoreParser.
type ScoreListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterScore is called when entering the score production.
	EnterScore(c *ScoreContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitScore is called when exiting the score production.
	ExitScore(c *ScoreContext)
}
