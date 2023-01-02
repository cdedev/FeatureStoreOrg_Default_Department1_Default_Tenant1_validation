// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672633185710/Z_Score.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Z_Score

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseZ_ScoreListener is a complete listener for a parse tree produced by Z_ScoreParser.
type BaseZ_ScoreListener struct{}

var _ Z_ScoreListener = &BaseZ_ScoreListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZ_ScoreListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZ_ScoreListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZ_ScoreListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZ_ScoreListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseZ_ScoreListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseZ_ScoreListener) ExitExpression(ctx *ExpressionContext) {}

// EnterZ_score is called when production z_score is entered.
func (s *BaseZ_ScoreListener) EnterZ_score(ctx *Z_scoreContext) {}

// ExitZ_score is called when production z_score is exited.
func (s *BaseZ_ScoreListener) ExitZ_score(ctx *Z_scoreContext) {}
