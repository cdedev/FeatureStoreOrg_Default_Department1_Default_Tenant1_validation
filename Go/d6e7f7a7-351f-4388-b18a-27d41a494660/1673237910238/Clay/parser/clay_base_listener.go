// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673237910238/Clay.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Clay

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseClayListener is a complete listener for a parse tree produced by ClayParser.
type BaseClayListener struct{}

var _ ClayListener = &BaseClayListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseClayListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseClayListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseClayListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseClayListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseClayListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseClayListener) ExitExpression(ctx *ExpressionContext) {}

// EnterClay is called when production clay is entered.
func (s *BaseClayListener) EnterClay(ctx *ClayContext) {}

// ExitClay is called when production clay is exited.
func (s *BaseClayListener) ExitClay(ctx *ClayContext) {}
