// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675230587509/Episode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Episode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEpisodeListener is a complete listener for a parse tree produced by EpisodeParser.
type BaseEpisodeListener struct{}

var _ EpisodeListener = &BaseEpisodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEpisodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEpisodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEpisodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEpisodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEpisodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEpisodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEpisode is called when production episode is entered.
func (s *BaseEpisodeListener) EnterEpisode(ctx *EpisodeContext) {}

// ExitEpisode is called when production episode is exited.
func (s *BaseEpisodeListener) ExitEpisode(ctx *EpisodeContext) {}
