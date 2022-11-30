// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780211583/Capacity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Capacity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCapacityListener is a complete listener for a parse tree produced by CapacityParser.
type BaseCapacityListener struct{}

var _ CapacityListener = &BaseCapacityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCapacityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCapacityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCapacityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCapacityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCapacityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCapacityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCapacity is called when production capacity is entered.
func (s *BaseCapacityListener) EnterCapacity(ctx *CapacityContext) {}

// ExitCapacity is called when production capacity is exited.
func (s *BaseCapacityListener) ExitCapacity(ctx *CapacityContext) {}
