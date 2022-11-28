// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669657111128/Wind.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wind

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWindListener is a complete listener for a parse tree produced by WindParser.
type BaseWindListener struct{}

var _ WindListener = &BaseWindListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWindListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWindListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWindListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWindListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWindListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWindListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWind is called when production wind is entered.
func (s *BaseWindListener) EnterWind(ctx *WindContext) {}

// ExitWind is called when production wind is exited.
func (s *BaseWindListener) ExitWind(ctx *WindContext) {}
