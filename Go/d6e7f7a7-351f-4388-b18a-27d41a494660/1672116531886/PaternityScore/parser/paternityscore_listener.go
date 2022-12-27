// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116531886/PaternityScore.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PaternityScore

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PaternityScoreListener is a complete listener for a parse tree produced by PaternityScoreParser.
type PaternityScoreListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPaternityscore is called when entering the paternityscore production.
	EnterPaternityscore(c *PaternityscoreContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPaternityscore is called when exiting the paternityscore production.
	ExitPaternityscore(c *PaternityscoreContext)
}
