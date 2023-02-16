// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522166037/Sick.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sick

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SickListener is a complete listener for a parse tree produced by SickParser.
type SickListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSick is called when entering the sick production.
	EnterSick(c *SickContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSick is called when exiting the sick production.
	ExitSick(c *SickContext)
}
