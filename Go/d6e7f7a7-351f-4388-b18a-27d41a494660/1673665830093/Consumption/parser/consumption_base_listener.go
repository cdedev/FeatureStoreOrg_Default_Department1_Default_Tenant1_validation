// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673665830093/Consumption.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Consumption

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConsumptionListener is a complete listener for a parse tree produced by ConsumptionParser.
type BaseConsumptionListener struct{}

var _ ConsumptionListener = &BaseConsumptionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConsumptionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConsumptionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConsumptionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConsumptionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConsumptionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConsumptionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConsumption is called when production consumption is entered.
func (s *BaseConsumptionListener) EnterConsumption(ctx *ConsumptionContext) {}

// ExitConsumption is called when production consumption is exited.
func (s *BaseConsumptionListener) ExitConsumption(ctx *ConsumptionContext) {}
