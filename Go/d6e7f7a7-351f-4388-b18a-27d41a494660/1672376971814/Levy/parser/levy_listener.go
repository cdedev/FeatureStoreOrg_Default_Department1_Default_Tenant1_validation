// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376971814/Levy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Levy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LevyListener is a complete listener for a parse tree produced by LevyParser.
type LevyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLevy is called when entering the levy production.
	EnterLevy(c *LevyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLevy is called when exiting the levy production.
	ExitLevy(c *LevyContext)
}
