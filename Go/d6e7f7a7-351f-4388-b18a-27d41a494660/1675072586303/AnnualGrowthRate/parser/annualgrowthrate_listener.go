// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072586303/AnnualGrowthRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AnnualGrowthRate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AnnualGrowthRateListener is a complete listener for a parse tree produced by AnnualGrowthRateParser.
type AnnualGrowthRateListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAnnualgrowthrate is called when entering the annualgrowthrate production.
	EnterAnnualgrowthrate(c *AnnualgrowthrateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAnnualgrowthrate is called when exiting the annualgrowthrate production.
	ExitAnnualgrowthrate(c *AnnualgrowthrateContext)
}
