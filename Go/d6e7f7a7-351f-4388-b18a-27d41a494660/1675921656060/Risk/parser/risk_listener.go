// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675921656060/Risk.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Risk

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RiskListener is a complete listener for a parse tree produced by RiskParser.
type RiskListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRisk is called when entering the risk production.
	EnterRisk(c *RiskContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRisk is called when exiting the risk production.
	ExitRisk(c *RiskContext)
}
