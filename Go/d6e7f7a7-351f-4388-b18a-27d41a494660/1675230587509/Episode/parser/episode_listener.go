// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675230587509/Episode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Episode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// EpisodeListener is a complete listener for a parse tree produced by EpisodeParser.
type EpisodeListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterEpisode is called when entering the episode production.
	EnterEpisode(c *EpisodeContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitEpisode is called when exiting the episode production.
	ExitEpisode(c *EpisodeContext)
}
