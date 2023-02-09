// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675920663558/Label.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Label

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LabelListener is a complete listener for a parse tree produced by LabelParser.
type LabelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)
}
