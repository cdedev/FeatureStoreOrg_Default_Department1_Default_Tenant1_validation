// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673929021260/Status.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Status

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStatusListener is a complete listener for a parse tree produced by StatusParser.
type BaseStatusListener struct{}

var _ StatusListener = &BaseStatusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStatusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStatusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStatusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStatusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStatusListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStatusListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStatus is called when production status is entered.
func (s *BaseStatusListener) EnterStatus(ctx *StatusContext) {}

// ExitStatus is called when production status is exited.
func (s *BaseStatusListener) ExitStatus(ctx *StatusContext) {}
