// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674538930547/Trails.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Trails

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTrailsListener is a complete listener for a parse tree produced by TrailsParser.
type BaseTrailsListener struct{}

var _ TrailsListener = &BaseTrailsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTrailsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTrailsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTrailsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTrailsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTrailsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTrailsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTrails is called when production trails is entered.
func (s *BaseTrailsListener) EnterTrails(ctx *TrailsContext) {}

// ExitTrails is called when production trails is exited.
func (s *BaseTrailsListener) ExitTrails(ctx *TrailsContext) {}
