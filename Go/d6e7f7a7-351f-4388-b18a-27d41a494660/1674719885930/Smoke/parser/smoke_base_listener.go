// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674719885930/Smoke.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Smoke

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSmokeListener is a complete listener for a parse tree produced by SmokeParser.
type BaseSmokeListener struct{}

var _ SmokeListener = &BaseSmokeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSmokeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSmokeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSmokeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSmokeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSmokeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSmokeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSmoke is called when production smoke is entered.
func (s *BaseSmokeListener) EnterSmoke(ctx *SmokeContext) {}

// ExitSmoke is called when production smoke is exited.
func (s *BaseSmokeListener) ExitSmoke(ctx *SmokeContext) {}
