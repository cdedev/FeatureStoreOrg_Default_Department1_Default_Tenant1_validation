// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674883414607/TotalLikes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TotalLikes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TotalLikesListener is a complete listener for a parse tree produced by TotalLikesParser.
type TotalLikesListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterTotallikes is called when entering the totallikes production.
	EnterTotallikes(c *TotallikesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitTotallikes is called when exiting the totallikes production.
	ExitTotallikes(c *TotallikesContext)
}
