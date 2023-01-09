// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673236162969/Survived.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Survived

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SurvivedListener is a complete listener for a parse tree produced by SurvivedParser.
type SurvivedListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSurvived is called when entering the survived production.
	EnterSurvived(c *SurvivedContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSurvived is called when exiting the survived production.
	ExitSurvived(c *SurvivedContext)
}
