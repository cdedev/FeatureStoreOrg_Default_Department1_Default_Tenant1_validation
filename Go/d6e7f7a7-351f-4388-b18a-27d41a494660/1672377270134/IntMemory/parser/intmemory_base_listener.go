// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377270134/IntMemory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // IntMemory

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIntMemoryListener is a complete listener for a parse tree produced by IntMemoryParser.
type BaseIntMemoryListener struct{}

var _ IntMemoryListener = &BaseIntMemoryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIntMemoryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIntMemoryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIntMemoryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIntMemoryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIntMemoryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIntMemoryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIntmemory is called when production intmemory is entered.
func (s *BaseIntMemoryListener) EnterIntmemory(ctx *IntmemoryContext) {}

// ExitIntmemory is called when production intmemory is exited.
func (s *BaseIntMemoryListener) ExitIntmemory(ctx *IntmemoryContext) {}
