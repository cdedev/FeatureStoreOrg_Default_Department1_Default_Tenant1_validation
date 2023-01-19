// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096667802/Gain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gain

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGainListener is a complete listener for a parse tree produced by GainParser.
type BaseGainListener struct{}

var _ GainListener = &BaseGainListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGainListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGainListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGainListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGainListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGainListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGainListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGain is called when production gain is entered.
func (s *BaseGainListener) EnterGain(ctx *GainContext) {}

// ExitGain is called when production gain is exited.
func (s *BaseGainListener) ExitGain(ctx *GainContext) {}
