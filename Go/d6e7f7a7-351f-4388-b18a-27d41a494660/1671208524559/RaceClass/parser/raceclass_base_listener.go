// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671208524559/RaceClass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RaceClass

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRaceClassListener is a complete listener for a parse tree produced by RaceClassParser.
type BaseRaceClassListener struct{}

var _ RaceClassListener = &BaseRaceClassListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRaceClassListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRaceClassListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRaceClassListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRaceClassListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRaceClassListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRaceClassListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRaceclass is called when production raceclass is entered.
func (s *BaseRaceClassListener) EnterRaceclass(ctx *RaceclassContext) {}

// ExitRaceclass is called when production raceclass is exited.
func (s *BaseRaceClassListener) ExitRaceclass(ctx *RaceclassContext) {}
