// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675240859030/Penality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Penality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePenalityListener is a complete listener for a parse tree produced by PenalityParser.
type BasePenalityListener struct{}

var _ PenalityListener = &BasePenalityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePenalityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePenalityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePenalityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePenalityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePenalityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePenalityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPenality is called when production penality is entered.
func (s *BasePenalityListener) EnterPenality(ctx *PenalityContext) {}

// ExitPenality is called when production penality is exited.
func (s *BasePenalityListener) ExitPenality(ctx *PenalityContext) {}
