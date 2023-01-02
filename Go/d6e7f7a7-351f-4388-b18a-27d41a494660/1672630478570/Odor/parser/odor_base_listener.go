// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672630478570/Odor.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Odor

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOdorListener is a complete listener for a parse tree produced by OdorParser.
type BaseOdorListener struct{}

var _ OdorListener = &BaseOdorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOdorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOdorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOdorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOdorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOdorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOdorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOdor is called when production odor is entered.
func (s *BaseOdorListener) EnterOdor(ctx *OdorContext) {}

// ExitOdor is called when production odor is exited.
func (s *BaseOdorListener) ExitOdor(ctx *OdorContext) {}
