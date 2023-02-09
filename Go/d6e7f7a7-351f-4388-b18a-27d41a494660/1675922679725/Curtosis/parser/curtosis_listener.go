// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675922679725/Curtosis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Curtosis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CurtosisListener is a complete listener for a parse tree produced by CurtosisParser.
type CurtosisListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCurtosis is called when entering the curtosis production.
	EnterCurtosis(c *CurtosisContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCurtosis is called when exiting the curtosis production.
	ExitCurtosis(c *CurtosisContext)
}
