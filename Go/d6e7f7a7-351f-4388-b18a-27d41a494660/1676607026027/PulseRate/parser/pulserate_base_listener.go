// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676607026027/PulseRate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PulseRate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePulseRateListener is a complete listener for a parse tree produced by PulseRateParser.
type BasePulseRateListener struct{}

var _ PulseRateListener = &BasePulseRateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePulseRateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePulseRateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePulseRateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePulseRateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePulseRateListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePulseRateListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPulserate is called when production pulserate is entered.
func (s *BasePulseRateListener) EnterPulserate(ctx *PulserateContext) {}

// ExitPulserate is called when production pulserate is exited.
func (s *BasePulseRateListener) ExitPulserate(ctx *PulserateContext) {}
