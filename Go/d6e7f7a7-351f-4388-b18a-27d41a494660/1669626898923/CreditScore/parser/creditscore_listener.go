// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669626898923/CreditScore.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CreditScore

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CreditScoreListener is a complete listener for a parse tree produced by CreditScoreParser.
type CreditScoreListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCreditscore is called when entering the creditscore production.
	EnterCreditscore(c *CreditscoreContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCreditscore is called when exiting the creditscore production.
	ExitCreditscore(c *CreditscoreContext)
}
