// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235280067/Eonal_stability.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Eonal_stability

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Eonal_stabilityListener is a complete listener for a parse tree produced by Eonal_stabilityParser.
type Eonal_stabilityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEonal_stability is called when entering the eonal_stability production.
	EnterEonal_stability(c *Eonal_stabilityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEonal_stability is called when exiting the eonal_stability production.
	ExitEonal_stability(c *Eonal_stabilityContext)
}
