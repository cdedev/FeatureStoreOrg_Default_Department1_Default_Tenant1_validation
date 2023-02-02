// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675315075632/Federation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Federation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FederationListener is a complete listener for a parse tree produced by FederationParser.
type FederationListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFederation is called when entering the federation production.
	EnterFederation(c *FederationContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFederation is called when exiting the federation production.
	ExitFederation(c *FederationContext)
}
