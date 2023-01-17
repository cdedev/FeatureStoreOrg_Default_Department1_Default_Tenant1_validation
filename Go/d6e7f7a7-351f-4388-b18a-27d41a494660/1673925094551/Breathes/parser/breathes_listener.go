// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925094551/Breathes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Breathes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BreathesListener is a complete listener for a parse tree produced by BreathesParser.
type BreathesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBreathes is called when entering the breathes production.
	EnterBreathes(c *BreathesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBreathes is called when exiting the breathes production.
	ExitBreathes(c *BreathesContext)
}
