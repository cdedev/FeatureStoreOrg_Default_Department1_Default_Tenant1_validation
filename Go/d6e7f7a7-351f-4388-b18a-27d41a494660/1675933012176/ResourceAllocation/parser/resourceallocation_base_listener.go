// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675933012176/ResourceAllocation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ResourceAllocation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseResourceAllocationListener is a complete listener for a parse tree produced by ResourceAllocationParser.
type BaseResourceAllocationListener struct{}

var _ ResourceAllocationListener = &BaseResourceAllocationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseResourceAllocationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseResourceAllocationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseResourceAllocationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseResourceAllocationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseResourceAllocationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseResourceAllocationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterResourceallocation is called when production resourceallocation is entered.
func (s *BaseResourceAllocationListener) EnterResourceallocation(ctx *ResourceallocationContext) {}

// ExitResourceallocation is called when production resourceallocation is exited.
func (s *BaseResourceAllocationListener) ExitResourceallocation(ctx *ResourceallocationContext) {}
