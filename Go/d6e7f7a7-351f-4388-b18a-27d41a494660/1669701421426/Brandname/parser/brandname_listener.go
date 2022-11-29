// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669701421426/Brandname.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Brandname

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BrandnameListener is a complete listener for a parse tree produced by BrandnameParser.
type BrandnameListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBrandname is called when entering the brandname production.
	EnterBrandname(c *BrandnameContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBrandname is called when exiting the brandname production.
	ExitBrandname(c *BrandnameContext)
}
