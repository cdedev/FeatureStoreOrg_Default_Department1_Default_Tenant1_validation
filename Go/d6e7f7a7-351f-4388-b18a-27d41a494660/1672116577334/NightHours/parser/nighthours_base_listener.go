// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672116577334/NightHours.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NightHours

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNightHoursListener is a complete listener for a parse tree produced by NightHoursParser.
type BaseNightHoursListener struct{}

var _ NightHoursListener = &BaseNightHoursListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNightHoursListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNightHoursListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNightHoursListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNightHoursListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNightHoursListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNightHoursListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNighthours is called when production nighthours is entered.
func (s *BaseNightHoursListener) EnterNighthours(ctx *NighthoursContext) {}

// ExitNighthours is called when production nighthours is exited.
func (s *BaseNightHoursListener) ExitNighthours(ctx *NighthoursContext) {}
