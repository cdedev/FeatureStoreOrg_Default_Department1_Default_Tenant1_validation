// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675845518714/Constituency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Constituency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ConstituencyListener is a complete listener for a parse tree produced by ConstituencyParser.
type ConstituencyListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstituency is called when entering the constituency production.
	EnterConstituency(c *ConstituencyContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstituency is called when exiting the constituency production.
	ExitConstituency(c *ConstituencyContext)
}
