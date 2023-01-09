// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673235389726/Extraversion.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Extraversion

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExtraversionListener is a complete listener for a parse tree produced by ExtraversionParser.
type ExtraversionListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExtraversion is called when entering the extraversion production.
	EnterExtraversion(c *ExtraversionContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExtraversion is called when exiting the extraversion production.
	ExitExtraversion(c *ExtraversionContext)
}
