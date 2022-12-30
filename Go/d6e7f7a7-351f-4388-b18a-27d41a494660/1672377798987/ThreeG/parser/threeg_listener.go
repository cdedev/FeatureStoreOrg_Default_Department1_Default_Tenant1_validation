// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377798987/ThreeG.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ThreeG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ThreeGListener is a complete listener for a parse tree produced by ThreeGParser.
type ThreeGListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterThreeg is called when entering the threeg production.
	EnterThreeg(c *ThreegContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitThreeg is called when exiting the threeg production.
	ExitThreeg(c *ThreegContext)
}
