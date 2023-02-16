// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676517712164/Tenure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tenure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TenureListener is a complete listener for a parse tree produced by TenureParser.
type TenureListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTenure is called when entering the tenure production.
	EnterTenure(c *TenureContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTenure is called when exiting the tenure production.
	ExitTenure(c *TenureContext)
}
