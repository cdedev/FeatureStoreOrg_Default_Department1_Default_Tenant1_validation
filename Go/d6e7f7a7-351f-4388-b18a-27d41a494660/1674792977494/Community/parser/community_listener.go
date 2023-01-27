// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674792977494/Community.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Community

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CommunityListener is a complete listener for a parse tree produced by CommunityParser.
type CommunityListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterCommunity is called when entering the community production.
	EnterCommunity(c *CommunityContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitCommunity is called when exiting the community production.
	ExitCommunity(c *CommunityContext)
}
