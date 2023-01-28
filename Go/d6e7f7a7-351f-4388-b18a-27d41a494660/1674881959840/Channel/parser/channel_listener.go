// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674881959840/Channel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Channel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChannelListener is a complete listener for a parse tree produced by ChannelParser.
type ChannelListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterChannel is called when entering the channel production.
	EnterChannel(c *ChannelContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitChannel is called when exiting the channel production.
	ExitChannel(c *ChannelContext)
}
