// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675149663663/Route.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Route

import "github.com/antlr/antlr4/runtime/Go/antlr"

// RouteListener is a complete listener for a parse tree produced by RouteParser.
type RouteListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterRoute is called when entering the route production.
	EnterRoute(c *RouteContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitRoute is called when exiting the route production.
	ExitRoute(c *RouteContext)
}
