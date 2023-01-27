// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674792391904/Waterproof.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Waterproof

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWaterproofListener is a complete listener for a parse tree produced by WaterproofParser.
type BaseWaterproofListener struct{}

var _ WaterproofListener = &BaseWaterproofListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWaterproofListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWaterproofListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWaterproofListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWaterproofListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWaterproofListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWaterproofListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWaterproof is called when production waterproof is entered.
func (s *BaseWaterproofListener) EnterWaterproof(ctx *WaterproofContext) {}

// ExitWaterproof is called when production waterproof is exited.
func (s *BaseWaterproofListener) ExitWaterproof(ctx *WaterproofContext) {}
