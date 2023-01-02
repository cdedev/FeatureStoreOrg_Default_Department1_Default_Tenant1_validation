// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672633185710/Z_Score.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Z_Score

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Z_ScoreListener is a complete listener for a parse tree produced by Z_ScoreParser.
type Z_ScoreListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterZ_score is called when entering the z_score production.
	EnterZ_score(c *Z_scoreContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitZ_score is called when exiting the z_score production.
	ExitZ_score(c *Z_scoreContext)
}
