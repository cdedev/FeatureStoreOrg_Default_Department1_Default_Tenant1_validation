// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672197300625/Traffic_Density_Score.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Traffic_Density_Score

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Traffic_Density_ScoreListener is a complete listener for a parse tree produced by Traffic_Density_ScoreParser.
type Traffic_Density_ScoreListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTraffic_density_score is called when entering the traffic_density_score production.
	EnterTraffic_density_score(c *Traffic_density_scoreContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTraffic_density_score is called when exiting the traffic_density_score production.
	ExitTraffic_density_score(c *Traffic_density_scoreContext)
}
