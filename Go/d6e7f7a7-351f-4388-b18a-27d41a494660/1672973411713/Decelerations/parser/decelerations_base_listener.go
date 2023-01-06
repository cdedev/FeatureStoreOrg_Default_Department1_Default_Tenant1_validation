// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672973411713/Decelerations.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Decelerations

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDecelerationsListener is a complete listener for a parse tree produced by DecelerationsParser.
type BaseDecelerationsListener struct{}

var _ DecelerationsListener = &BaseDecelerationsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDecelerationsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDecelerationsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDecelerationsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDecelerationsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDecelerationsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDecelerationsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDecelerations is called when production decelerations is entered.
func (s *BaseDecelerationsListener) EnterDecelerations(ctx *DecelerationsContext) {}

// ExitDecelerations is called when production decelerations is exited.
func (s *BaseDecelerationsListener) ExitDecelerations(ctx *DecelerationsContext) {}
