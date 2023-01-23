// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447529434/Visit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Visit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// VisitListener is a complete listener for a parse tree produced by VisitParser.
type VisitListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterVisit is called when entering the visit production.
	EnterVisit(c *VisitContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitVisit is called when exiting the visit production.
	ExitVisit(c *VisitContext)
}
