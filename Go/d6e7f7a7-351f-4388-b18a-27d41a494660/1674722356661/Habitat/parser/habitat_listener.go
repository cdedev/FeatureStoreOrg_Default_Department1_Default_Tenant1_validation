// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674722356661/Habitat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Habitat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HabitatListener is a complete listener for a parse tree produced by HabitatParser.
type HabitatListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHabitat is called when entering the habitat production.
	EnterHabitat(c *HabitatContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHabitat is called when exiting the habitat production.
	ExitHabitat(c *HabitatContext)
}
