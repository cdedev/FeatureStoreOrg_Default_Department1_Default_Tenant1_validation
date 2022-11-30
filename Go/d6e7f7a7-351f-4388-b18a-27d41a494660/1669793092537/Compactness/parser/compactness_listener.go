// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669793092537/Compactness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Compactness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CompactnessListener is a complete listener for a parse tree produced by CompactnessParser.
type CompactnessListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCompactness is called when entering the compactness production.
	EnterCompactness(c *CompactnessContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCompactness is called when exiting the compactness production.
	ExitCompactness(c *CompactnessContext)
}
