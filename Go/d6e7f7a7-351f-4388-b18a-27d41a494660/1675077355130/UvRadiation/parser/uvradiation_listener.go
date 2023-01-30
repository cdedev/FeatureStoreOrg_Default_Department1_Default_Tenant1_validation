// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675077355130/UvRadiation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // UvRadiation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// UvRadiationListener is a complete listener for a parse tree produced by UvRadiationParser.
type UvRadiationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterUvradiation is called when entering the uvradiation production.
	EnterUvradiation(c *UvradiationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitUvradiation is called when exiting the uvradiation production.
	ExitUvradiation(c *UvradiationContext)
}
