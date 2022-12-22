// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671689605713/Centroid.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Centroid

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CentroidListener is a complete listener for a parse tree produced by CentroidParser.
type CentroidListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCentroid is called when entering the centroid production.
	EnterCentroid(c *CentroidContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCentroid is called when exiting the centroid production.
	ExitCentroid(c *CentroidContext)
}
