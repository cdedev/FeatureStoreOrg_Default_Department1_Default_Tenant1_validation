// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675834897679/Posture.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Posture

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PostureListener is a complete listener for a parse tree produced by PostureParser.
type PostureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPosture is called when entering the posture production.
	EnterPosture(c *PostureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPosture is called when exiting the posture production.
	ExitPosture(c *PostureContext)
}
