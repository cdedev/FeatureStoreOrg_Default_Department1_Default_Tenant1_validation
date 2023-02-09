// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675932910272/Designation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Designation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// DesignationListener is a complete listener for a parse tree produced by DesignationParser.
type DesignationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterDesignation is called when entering the designation production.
	EnterDesignation(c *DesignationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitDesignation is called when exiting the designation production.
	ExitDesignation(c *DesignationContext)
}
