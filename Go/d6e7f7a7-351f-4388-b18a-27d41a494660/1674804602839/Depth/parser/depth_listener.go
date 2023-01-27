// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674804602839/Depth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Depth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DepthListener is a complete listener for a parse tree produced by DepthParser.
type DepthListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDepth is called when entering the depth production.
	EnterDepth(c *DepthContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDepth is called when exiting the depth production.
	ExitDepth(c *DepthContext)
}
