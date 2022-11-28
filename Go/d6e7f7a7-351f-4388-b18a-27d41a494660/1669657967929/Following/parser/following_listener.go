// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669657967929/Following.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Following

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FollowingListener is a complete listener for a parse tree produced by FollowingParser.
type FollowingListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterFollowing is called when entering the following production.
	EnterFollowing(c *FollowingContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitFollowing is called when exiting the following production.
	ExitFollowing(c *FollowingContext)
}
