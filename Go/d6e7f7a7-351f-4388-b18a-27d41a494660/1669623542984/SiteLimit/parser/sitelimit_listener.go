// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623542984/SiteLimit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SiteLimit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SiteLimitListener is a complete listener for a parse tree produced by SiteLimitParser.
type SiteLimitListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterSitelimit is called when entering the sitelimit production.
	EnterSitelimit(c *SitelimitContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitSitelimit is called when exiting the sitelimit production.
	ExitSitelimit(c *SitelimitContext)
}
