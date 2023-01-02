// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672629823235/Tempo.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tempo

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTempoListener is a complete listener for a parse tree produced by TempoParser.
type BaseTempoListener struct{}

var _ TempoListener = &BaseTempoListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTempoListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTempoListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTempoListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTempoListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTempoListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTempoListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTempo is called when production tempo is entered.
func (s *BaseTempoListener) EnterTempo(ctx *TempoContext) {}

// ExitTempo is called when production tempo is exited.
func (s *BaseTempoListener) ExitTempo(ctx *TempoContext) {}
