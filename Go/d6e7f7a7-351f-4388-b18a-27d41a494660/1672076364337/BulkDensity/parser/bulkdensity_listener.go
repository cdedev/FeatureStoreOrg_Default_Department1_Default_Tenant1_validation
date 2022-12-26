// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076364337/BulkDensity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BulkDensity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BulkDensityListener is a complete listener for a parse tree produced by BulkDensityParser.
type BulkDensityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBulkdensity is called when entering the bulkdensity production.
	EnterBulkdensity(c *BulkdensityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBulkdensity is called when exiting the bulkdensity production.
	ExitBulkdensity(c *BulkdensityContext)
}
