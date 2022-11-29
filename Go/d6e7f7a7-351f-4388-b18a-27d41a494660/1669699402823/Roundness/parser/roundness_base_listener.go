// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699402823/Roundness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Roundness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRoundnessListener is a complete listener for a parse tree produced by RoundnessParser.
type BaseRoundnessListener struct{}

var _ RoundnessListener = &BaseRoundnessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRoundnessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRoundnessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRoundnessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRoundnessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRoundnessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRoundnessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRoundness is called when production roundness is entered.
func (s *BaseRoundnessListener) EnterRoundness(ctx *RoundnessContext) {}

// ExitRoundness is called when production roundness is exited.
func (s *BaseRoundnessListener) ExitRoundness(ctx *RoundnessContext) {}
