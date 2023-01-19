// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674104270059/GlobalRadiation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GlobalRadiation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GlobalRadiationListener is a complete listener for a parse tree produced by GlobalRadiationParser.
type GlobalRadiationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGlobalradiation is called when entering the globalradiation production.
	EnterGlobalradiation(c *GlobalradiationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGlobalradiation is called when exiting the globalradiation production.
	ExitGlobalradiation(c *GlobalradiationContext)
}
