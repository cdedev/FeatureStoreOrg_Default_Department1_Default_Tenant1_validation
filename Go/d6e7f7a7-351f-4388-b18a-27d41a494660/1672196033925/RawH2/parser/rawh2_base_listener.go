// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196033925/RawH2.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RawH2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRawH2Listener is a complete listener for a parse tree produced by RawH2Parser.
type BaseRawH2Listener struct{}

var _ RawH2Listener = &BaseRawH2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRawH2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRawH2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRawH2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRawH2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRawH2Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRawH2Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterRawh2 is called when production rawh2 is entered.
func (s *BaseRawH2Listener) EnterRawh2(ctx *Rawh2Context) {}

// ExitRawh2 is called when production rawh2 is exited.
func (s *BaseRawH2Listener) ExitRawh2(ctx *Rawh2Context) {}
