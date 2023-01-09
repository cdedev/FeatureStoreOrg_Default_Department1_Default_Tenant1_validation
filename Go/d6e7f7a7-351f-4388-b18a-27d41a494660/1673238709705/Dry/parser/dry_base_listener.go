// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238709705/Dry.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dry

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDryListener is a complete listener for a parse tree produced by DryParser.
type BaseDryListener struct{}

var _ DryListener = &BaseDryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDry is called when production dry is entered.
func (s *BaseDryListener) EnterDry(ctx *DryContext) {}

// ExitDry is called when production dry is exited.
func (s *BaseDryListener) ExitDry(ctx *DryContext) {}
