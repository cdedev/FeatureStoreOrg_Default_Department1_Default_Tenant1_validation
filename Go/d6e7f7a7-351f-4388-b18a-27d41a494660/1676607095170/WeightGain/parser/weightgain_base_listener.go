// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676607095170/WeightGain.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WeightGain

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWeightGainListener is a complete listener for a parse tree produced by WeightGainParser.
type BaseWeightGainListener struct{}

var _ WeightGainListener = &BaseWeightGainListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWeightGainListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWeightGainListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWeightGainListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWeightGainListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWeightGainListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWeightGainListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWeightgain is called when production weightgain is entered.
func (s *BaseWeightGainListener) EnterWeightgain(ctx *WeightgainContext) {}

// ExitWeightgain is called when production weightgain is exited.
func (s *BaseWeightGainListener) ExitWeightgain(ctx *WeightgainContext) {}
