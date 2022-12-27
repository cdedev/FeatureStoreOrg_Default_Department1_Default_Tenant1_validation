// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115062165/Traffic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Traffic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTrafficListener is a complete listener for a parse tree produced by TrafficParser.
type BaseTrafficListener struct{}

var _ TrafficListener = &BaseTrafficListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTrafficListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTrafficListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTrafficListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTrafficListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTrafficListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTrafficListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTraffic is called when production traffic is entered.
func (s *BaseTrafficListener) EnterTraffic(ctx *TrafficContext) {}

// ExitTraffic is called when production traffic is exited.
func (s *BaseTrafficListener) ExitTraffic(ctx *TrafficContext) {}
