// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778891384/Stream.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Stream

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStreamListener is a complete listener for a parse tree produced by StreamParser.
type BaseStreamListener struct{}

var _ StreamListener = &BaseStreamListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStreamListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStreamListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStreamListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStreamListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStreamListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStreamListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStream is called when production stream is entered.
func (s *BaseStreamListener) EnterStream(ctx *StreamContext) {}

// ExitStream is called when production stream is exited.
func (s *BaseStreamListener) ExitStream(ctx *StreamContext) {}
