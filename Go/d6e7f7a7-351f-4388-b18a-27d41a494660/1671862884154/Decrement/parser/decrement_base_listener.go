// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671862884154/Decrement.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Decrement

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDecrementListener is a complete listener for a parse tree produced by DecrementParser.
type BaseDecrementListener struct{}

var _ DecrementListener = &BaseDecrementListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDecrementListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDecrementListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDecrementListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDecrementListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDecrementListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDecrementListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDecrement is called when production decrement is entered.
func (s *BaseDecrementListener) EnterDecrement(ctx *DecrementContext) {}

// ExitDecrement is called when production decrement is exited.
func (s *BaseDecrementListener) ExitDecrement(ctx *DecrementContext) {}
