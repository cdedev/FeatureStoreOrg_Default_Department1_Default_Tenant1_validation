// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867803590/Metal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Metal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MetalListener is a complete listener for a parse tree produced by MetalParser.
type MetalListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterMetal is called when entering the metal production.
	EnterMetal(c *MetalContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitMetal is called when exiting the metal production.
	ExitMetal(c *MetalContext)
}
