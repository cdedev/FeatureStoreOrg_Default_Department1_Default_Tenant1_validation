// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674720957369/Benefits.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Benefits

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BenefitsListener is a complete listener for a parse tree produced by BenefitsParser.
type BenefitsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBenefits is called when entering the benefits production.
	EnterBenefits(c *BenefitsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBenefits is called when exiting the benefits production.
	ExitBenefits(c *BenefitsContext)
}
