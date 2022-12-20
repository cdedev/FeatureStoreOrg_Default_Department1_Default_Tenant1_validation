// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671523287895/Season.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Season

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSeasonListener is a complete listener for a parse tree produced by SeasonParser.
type BaseSeasonListener struct{}

var _ SeasonListener = &BaseSeasonListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSeasonListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSeasonListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSeasonListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSeasonListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSeasonListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSeasonListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSeason is called when production season is entered.
func (s *BaseSeasonListener) EnterSeason(ctx *SeasonContext) {}

// ExitSeason is called when production season is exited.
func (s *BaseSeasonListener) ExitSeason(ctx *SeasonContext) {}
