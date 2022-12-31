// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672460625972/Hemisphere.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hemisphere

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HemisphereListener is a complete listener for a parse tree produced by HemisphereParser.
type HemisphereListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHemisphere is called when entering the hemisphere production.
	EnterHemisphere(c *HemisphereContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHemisphere is called when exiting the hemisphere production.
	ExitHemisphere(c *HemisphereContext)
}
