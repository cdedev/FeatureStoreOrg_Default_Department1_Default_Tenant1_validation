// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675317863660/Memory.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Memory

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMemoryListener is a complete listener for a parse tree produced by MemoryParser.
type BaseMemoryListener struct{}

var _ MemoryListener = &BaseMemoryListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMemoryListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMemoryListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMemoryListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMemoryListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMemoryListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMemoryListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMemory is called when production memory is entered.
func (s *BaseMemoryListener) EnterMemory(ctx *MemoryContext) {}

// ExitMemory is called when production memory is exited.
func (s *BaseMemoryListener) ExitMemory(ctx *MemoryContext) {}
