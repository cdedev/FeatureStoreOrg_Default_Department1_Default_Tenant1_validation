// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671203783527/RapeLegacy.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RapeLegacy

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RapeLegacyListener is a complete listener for a parse tree produced by RapeLegacyParser.
type RapeLegacyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRapelegacy is called when entering the rapelegacy production.
	EnterRapelegacy(c *RapelegacyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRapelegacy is called when exiting the rapelegacy production.
	ExitRapelegacy(c *RapelegacyContext)
}
