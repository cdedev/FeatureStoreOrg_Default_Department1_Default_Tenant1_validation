// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674799725772/OperatingSystem.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OperatingSystem

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OperatingSystemListener is a complete listener for a parse tree produced by OperatingSystemParser.
type OperatingSystemListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOperatingsystem is called when entering the operatingsystem production.
	EnterOperatingsystem(c *OperatingsystemContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOperatingsystem is called when exiting the operatingsystem production.
	ExitOperatingsystem(c *OperatingsystemContext)
}
