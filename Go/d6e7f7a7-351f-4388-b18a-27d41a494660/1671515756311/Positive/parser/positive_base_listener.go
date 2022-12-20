// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515756311/Positive.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Positive

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePositiveListener is a complete listener for a parse tree produced by PositiveParser.
type BasePositiveListener struct{}

var _ PositiveListener = &BasePositiveListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePositiveListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePositiveListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePositiveListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePositiveListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePositiveListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePositiveListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPositive is called when production positive is entered.
func (s *BasePositiveListener) EnterPositive(ctx *PositiveContext) {}

// ExitPositive is called when production positive is exited.
func (s *BasePositiveListener) ExitPositive(ctx *PositiveContext) {}
