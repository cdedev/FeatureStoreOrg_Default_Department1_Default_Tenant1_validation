// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196147965/RawEthanol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RawEthanol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RawEthanolListener is a complete listener for a parse tree produced by RawEthanolParser.
type RawEthanolListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRawethanol is called when entering the rawethanol production.
	EnterRawethanol(c *RawethanolContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRawethanol is called when exiting the rawethanol production.
	ExitRawethanol(c *RawethanolContext)
}
