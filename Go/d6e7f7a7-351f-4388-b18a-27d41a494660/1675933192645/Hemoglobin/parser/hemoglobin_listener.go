// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675933192645/Hemoglobin.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hemoglobin

import "github.com/antlr/antlr4/runtime/Go/antlr"

// HemoglobinListener is a complete listener for a parse tree produced by HemoglobinParser.
type HemoglobinListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterHemoglobin is called when entering the hemoglobin production.
	EnterHemoglobin(c *HemoglobinContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitHemoglobin is called when exiting the hemoglobin production.
	ExitHemoglobin(c *HemoglobinContext)
}
