// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674722475152/Occurrence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Occurrence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OccurrenceListener is a complete listener for a parse tree produced by OccurrenceParser.
type OccurrenceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOccurrence is called when entering the occurrence production.
	EnterOccurrence(c *OccurrenceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOccurrence is called when exiting the occurrence production.
	ExitOccurrence(c *OccurrenceContext)
}
