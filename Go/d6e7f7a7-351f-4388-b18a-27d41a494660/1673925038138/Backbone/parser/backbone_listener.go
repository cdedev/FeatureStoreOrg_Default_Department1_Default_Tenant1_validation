// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925038138/Backbone.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Backbone

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BackboneListener is a complete listener for a parse tree produced by BackboneParser.
type BackboneListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBackbone is called when entering the backbone production.
	EnterBackbone(c *BackboneContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBackbone is called when exiting the backbone production.
	ExitBackbone(c *BackboneContext)
}
