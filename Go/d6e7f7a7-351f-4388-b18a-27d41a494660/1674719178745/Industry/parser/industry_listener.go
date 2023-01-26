// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719178745/Industry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Industry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// IndustryListener is a complete listener for a parse tree produced by IndustryParser.
type IndustryListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterIndustry is called when entering the industry production.
	EnterIndustry(c *IndustryContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitIndustry is called when exiting the industry production.
	ExitIndustry(c *IndustryContext)
}
