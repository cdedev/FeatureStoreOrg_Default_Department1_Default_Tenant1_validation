// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675072188168/Safety.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Safety

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSafetyListener is a complete listener for a parse tree produced by SafetyParser.
type BaseSafetyListener struct{}

var _ SafetyListener = &BaseSafetyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSafetyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSafetyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSafetyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSafetyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSafetyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSafetyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSafety is called when production safety is entered.
func (s *BaseSafetyListener) EnterSafety(ctx *SafetyContext) {}

// ExitSafety is called when production safety is exited.
func (s *BaseSafetyListener) ExitSafety(ctx *SafetyContext) {}
