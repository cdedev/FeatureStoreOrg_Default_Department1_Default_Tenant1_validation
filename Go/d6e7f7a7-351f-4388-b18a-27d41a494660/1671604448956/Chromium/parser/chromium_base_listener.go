// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671604448956/Chromium.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chromium

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChromiumListener is a complete listener for a parse tree produced by ChromiumParser.
type BaseChromiumListener struct{}

var _ ChromiumListener = &BaseChromiumListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChromiumListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChromiumListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChromiumListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChromiumListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChromiumListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChromiumListener) ExitExpression(ctx *ExpressionContext) {}

// EnterChromium is called when production chromium is entered.
func (s *BaseChromiumListener) EnterChromium(ctx *ChromiumContext) {}

// ExitChromium is called when production chromium is exited.
func (s *BaseChromiumListener) ExitChromium(ctx *ChromiumContext) {}
