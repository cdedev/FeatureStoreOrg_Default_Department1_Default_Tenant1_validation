// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669786630654/Variety.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Variety

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VarietyListener is a complete listener for a parse tree produced by VarietyParser.
type VarietyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVariety is called when entering the variety production.
	EnterVariety(c *VarietyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVariety is called when exiting the variety production.
	ExitVariety(c *VarietyContext)
}
