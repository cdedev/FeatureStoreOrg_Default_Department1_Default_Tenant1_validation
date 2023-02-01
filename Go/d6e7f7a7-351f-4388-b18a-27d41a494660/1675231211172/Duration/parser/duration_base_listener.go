// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675231211172/Duration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Duration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDurationListener is a complete listener for a parse tree produced by DurationParser.
type BaseDurationListener struct{}

var _ DurationListener = &BaseDurationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDurationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDurationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDurationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDurationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDurationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDurationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDuration is called when production duration is entered.
func (s *BaseDurationListener) EnterDuration(ctx *DurationContext) {}

// ExitDuration is called when production duration is exited.
func (s *BaseDurationListener) ExitDuration(ctx *DurationContext) {}
