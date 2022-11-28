// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669652531904/Timestamp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Timestamp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TimestampListener is a complete listener for a parse tree produced by TimestampParser.
type TimestampListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTimestamp is called when entering the timestamp production.
	EnterTimestamp(c *TimestampContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTimestamp is called when exiting the timestamp production.
	ExitTimestamp(c *TimestampContext)
}
