// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671771855533/SequenceLengthCount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SequenceLengthCount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SequenceLengthCountListener is a complete listener for a parse tree produced by SequenceLengthCountParser.
type SequenceLengthCountListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSequencelengthcount is called when entering the sequencelengthcount production.
	EnterSequencelengthcount(c *SequencelengthcountContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSequencelengthcount is called when exiting the sequencelengthcount production.
	ExitSequencelengthcount(c *SequencelengthcountContext)
}
