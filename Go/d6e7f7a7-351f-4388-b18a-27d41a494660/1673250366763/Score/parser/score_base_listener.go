// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673250366763/Score.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Score

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseScoreListener is a complete listener for a parse tree produced by ScoreParser.
type BaseScoreListener struct{}

var _ ScoreListener = &BaseScoreListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseScoreListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseScoreListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseScoreListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseScoreListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseScoreListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseScoreListener) ExitExpression(ctx *ExpressionContext) {}

// EnterScore is called when production score is entered.
func (s *BaseScoreListener) EnterScore(ctx *ScoreContext) {}

// ExitScore is called when production score is exited.
func (s *BaseScoreListener) ExitScore(ctx *ScoreContext) {}
