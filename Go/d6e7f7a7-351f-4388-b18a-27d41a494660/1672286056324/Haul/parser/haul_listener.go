// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672286056324/Haul.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Haul

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HaulListener is a complete listener for a parse tree produced by HaulParser.
type HaulListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHaul is called when entering the haul production.
	EnterHaul(c *HaulContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHaul is called when exiting the haul production.
	ExitHaul(c *HaulContext)
}
