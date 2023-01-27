// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674800089241/Performance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Performance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// PerformanceListener is a complete listener for a parse tree produced by PerformanceParser.
type PerformanceListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPerformance is called when entering the performance production.
	EnterPerformance(c *PerformanceContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPerformance is called when exiting the performance production.
	ExitPerformance(c *PerformanceContext)
}
