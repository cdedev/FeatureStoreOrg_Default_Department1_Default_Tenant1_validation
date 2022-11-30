// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778934158/OnlineSecurity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OnlineSecurity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// OnlineSecurityListener is a complete listener for a parse tree produced by OnlineSecurityParser.
type OnlineSecurityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterOnlinesecurity is called when entering the onlinesecurity production.
	EnterOnlinesecurity(c *OnlinesecurityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitOnlinesecurity is called when exiting the onlinesecurity production.
	ExitOnlinesecurity(c *OnlinesecurityContext)
}
