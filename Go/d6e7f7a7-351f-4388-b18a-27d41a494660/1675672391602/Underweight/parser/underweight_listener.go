// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675672391602/Underweight.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Underweight

import "github.com/antlr/antlr4/runtime/Go/antlr"

// UnderweightListener is a complete listener for a parse tree produced by UnderweightParser.
type UnderweightListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUnderweight is called when entering the underweight production.
	EnterUnderweight(c *UnderweightContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUnderweight is called when exiting the underweight production.
	ExitUnderweight(c *UnderweightContext)
}
