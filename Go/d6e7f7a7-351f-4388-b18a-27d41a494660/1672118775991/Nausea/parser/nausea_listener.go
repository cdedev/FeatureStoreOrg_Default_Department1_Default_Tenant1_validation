// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118775991/Nausea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nausea

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NauseaListener is a complete listener for a parse tree produced by NauseaParser.
type NauseaListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNausea is called when entering the nausea production.
	EnterNausea(c *NauseaContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNausea is called when exiting the nausea production.
	ExitNausea(c *NauseaContext)
}
