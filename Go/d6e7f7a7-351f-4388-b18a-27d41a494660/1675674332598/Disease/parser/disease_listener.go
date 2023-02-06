// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675674332598/Disease.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Disease

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DiseaseListener is a complete listener for a parse tree produced by DiseaseParser.
type DiseaseListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDisease is called when entering the disease production.
	EnterDisease(c *DiseaseContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDisease is called when exiting the disease production.
	ExitDisease(c *DiseaseContext)
}
