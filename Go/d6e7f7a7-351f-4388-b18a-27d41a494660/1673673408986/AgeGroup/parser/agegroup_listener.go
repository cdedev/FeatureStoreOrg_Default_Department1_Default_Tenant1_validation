// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673673408986/AgeGroup.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AgeGroup

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AgeGroupListener is a complete listener for a parse tree produced by AgeGroupParser.
type AgeGroupListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAgegroup is called when entering the agegroup production.
	EnterAgegroup(c *AgegroupContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAgegroup is called when exiting the agegroup production.
	ExitAgegroup(c *AgegroupContext)
}
