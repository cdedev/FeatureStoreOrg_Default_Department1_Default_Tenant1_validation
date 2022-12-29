// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284999727/Lyrics.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lyrics

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLyricsListener is a complete listener for a parse tree produced by LyricsParser.
type BaseLyricsListener struct{}

var _ LyricsListener = &BaseLyricsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLyricsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLyricsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLyricsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLyricsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLyricsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLyricsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLyrics is called when production lyrics is entered.
func (s *BaseLyricsListener) EnterLyrics(ctx *LyricsContext) {}

// ExitLyrics is called when production lyrics is exited.
func (s *BaseLyricsListener) ExitLyrics(ctx *LyricsContext) {}
