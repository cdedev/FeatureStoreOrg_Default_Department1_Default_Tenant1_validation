// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284999727/Lyrics.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lyrics

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LyricsListener is a complete listener for a parse tree produced by LyricsParser.
type LyricsListener interface {
	antlr.ParseTreeListener

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLyrics is called when entering the lyrics production.
	EnterLyrics(c *LyricsContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLyrics is called when exiting the lyrics production.
	ExitLyrics(c *LyricsContext)
}
