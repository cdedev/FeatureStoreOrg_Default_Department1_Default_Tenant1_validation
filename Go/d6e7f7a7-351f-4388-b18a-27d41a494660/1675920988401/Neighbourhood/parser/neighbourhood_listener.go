// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675920988401/Neighbourhood.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Neighbourhood

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NeighbourhoodListener is a complete listener for a parse tree produced by NeighbourhoodParser.
type NeighbourhoodListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNeighbourhood is called when entering the neighbourhood production.
	EnterNeighbourhood(c *NeighbourhoodContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNeighbourhood is called when exiting the neighbourhood production.
	ExitNeighbourhood(c *NeighbourhoodContext)
}
