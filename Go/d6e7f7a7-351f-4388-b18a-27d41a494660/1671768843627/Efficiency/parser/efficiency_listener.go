// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671768843627/Efficiency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Efficiency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EfficiencyListener is a complete listener for a parse tree produced by EfficiencyParser.
type EfficiencyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEfficiency is called when entering the efficiency production.
	EnterEfficiency(c *EfficiencyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEfficiency is called when exiting the efficiency production.
	ExitEfficiency(c *EfficiencyContext)
}
