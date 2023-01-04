// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672802581911/WindDirection.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // WindDirection

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWindDirectionListener is a complete listener for a parse tree produced by WindDirectionParser.
type BaseWindDirectionListener struct{}

var _ WindDirectionListener = &BaseWindDirectionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWindDirectionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWindDirectionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWindDirectionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWindDirectionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWindDirectionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWindDirectionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWinddirection is called when production winddirection is entered.
func (s *BaseWindDirectionListener) EnterWinddirection(ctx *WinddirectionContext) {}

// ExitWinddirection is called when production winddirection is exited.
func (s *BaseWindDirectionListener) ExitWinddirection(ctx *WinddirectionContext) {}
