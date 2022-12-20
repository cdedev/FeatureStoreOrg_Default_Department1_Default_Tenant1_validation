// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671525209301/GDP.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GDP

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GDPListener is a complete listener for a parse tree produced by GDPParser.
type GDPListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterGdp is called when entering the gdp production.
	EnterGdp(c *GdpContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitGdp is called when exiting the gdp production.
	ExitGdp(c *GdpContext)
}
