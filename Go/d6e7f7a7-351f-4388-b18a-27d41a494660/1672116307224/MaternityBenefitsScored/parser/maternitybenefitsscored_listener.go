// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116307224/MaternityBenefitsScored.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MaternityBenefitsScored

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MaternityBenefitsScoredListener is a complete listener for a parse tree produced by MaternityBenefitsScoredParser.
type MaternityBenefitsScoredListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMaternitybenefitsscored is called when entering the maternitybenefitsscored production.
	EnterMaternitybenefitsscored(c *MaternitybenefitsscoredContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMaternitybenefitsscored is called when exiting the maternitybenefitsscored production.
	ExitMaternitybenefitsscored(c *MaternitybenefitsscoredContext)
}
