// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676524614179/InstrumentMode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InstrumentMode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInstrumentModeListener is a complete listener for a parse tree produced by InstrumentModeParser.
type BaseInstrumentModeListener struct{}

var _ InstrumentModeListener = &BaseInstrumentModeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInstrumentModeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInstrumentModeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInstrumentModeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInstrumentModeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInstrumentModeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInstrumentModeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInstrumentmode is called when production instrumentmode is entered.
func (s *BaseInstrumentModeListener) EnterInstrumentmode(ctx *InstrumentmodeContext) {}

// ExitInstrumentmode is called when production instrumentmode is exited.
func (s *BaseInstrumentModeListener) ExitInstrumentmode(ctx *InstrumentmodeContext) {}
