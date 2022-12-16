// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671207453503/Race.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Race

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRaceListener is a complete listener for a parse tree produced by RaceParser.
type BaseRaceListener struct{}

var _ RaceListener = &BaseRaceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRaceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRaceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRaceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRaceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRaceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRaceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRace is called when production race is entered.
func (s *BaseRaceListener) EnterRace(ctx *RaceContext) {}

// ExitRace is called when production race is exited.
func (s *BaseRaceListener) ExitRace(ctx *RaceContext) {}
