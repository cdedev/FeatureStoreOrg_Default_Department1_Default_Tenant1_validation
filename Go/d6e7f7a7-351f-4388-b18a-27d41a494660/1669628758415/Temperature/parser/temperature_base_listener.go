// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669628758415/Temperature.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Temperature

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTemperatureListener is a complete listener for a parse tree produced by TemperatureParser.
type BaseTemperatureListener struct{}

var _ TemperatureListener = &BaseTemperatureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTemperatureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTemperatureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTemperatureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTemperatureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTemperatureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTemperatureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTemperature is called when production temperature is entered.
func (s *BaseTemperatureListener) EnterTemperature(ctx *TemperatureContext) {}

// ExitTemperature is called when production temperature is exited.
func (s *BaseTemperatureListener) ExitTemperature(ctx *TemperatureContext) {}
