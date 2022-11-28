// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669633379954/Color.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Color

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseColorListener is a complete listener for a parse tree produced by ColorParser.
type BaseColorListener struct{}

var _ ColorListener = &BaseColorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseColorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseColorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseColorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseColorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseColorListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseColorListener) ExitExpression(ctx *ExpressionContext) {}

// EnterColor is called when production color is entered.
func (s *BaseColorListener) EnterColor(ctx *ColorContext) {}

// ExitColor is called when production color is exited.
func (s *BaseColorListener) ExitColor(ctx *ColorContext) {}
