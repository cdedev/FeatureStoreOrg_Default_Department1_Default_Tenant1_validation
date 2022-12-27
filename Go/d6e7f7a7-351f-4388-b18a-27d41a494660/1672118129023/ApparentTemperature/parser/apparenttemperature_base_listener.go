// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672118129023/ApparentTemperature.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ApparentTemperature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseApparentTemperatureListener is a complete listener for a parse tree produced by ApparentTemperatureParser.
type BaseApparentTemperatureListener struct{}

var _ ApparentTemperatureListener = &BaseApparentTemperatureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseApparentTemperatureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseApparentTemperatureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseApparentTemperatureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseApparentTemperatureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseApparentTemperatureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseApparentTemperatureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterApparenttemperature is called when production apparenttemperature is entered.
func (s *BaseApparentTemperatureListener) EnterApparenttemperature(ctx *ApparenttemperatureContext) {}

// ExitApparenttemperature is called when production apparenttemperature is exited.
func (s *BaseApparentTemperatureListener) ExitApparenttemperature(ctx *ApparenttemperatureContext) {}
