// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671862652685/LetterOfRecommendations.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LetterOfRecommendations

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LetterOfRecommendationsListener is a complete listener for a parse tree produced by LetterOfRecommendationsParser.
type LetterOfRecommendationsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLetterofrecommendations is called when entering the letterofrecommendations production.
	EnterLetterofrecommendations(c *LetterofrecommendationsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLetterofrecommendations is called when exiting the letterofrecommendations production.
	ExitLetterofrecommendations(c *LetterofrecommendationsContext)
}
