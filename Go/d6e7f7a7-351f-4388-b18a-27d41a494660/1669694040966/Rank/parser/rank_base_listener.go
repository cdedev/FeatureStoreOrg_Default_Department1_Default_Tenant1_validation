// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669694040966/Rank.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Rank

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRankListener is a complete listener for a parse tree produced by RankParser.
type BaseRankListener struct{}

var _ RankListener = &BaseRankListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRankListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRankListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRankListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRankListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRankListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRankListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRank is called when production rank is entered.
func (s *BaseRankListener) EnterRank(ctx *RankContext) {}

// ExitRank is called when production rank is exited.
func (s *BaseRankListener) ExitRank(ctx *RankContext) {}
