// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675068198075/Cpu.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cpu

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCpuListener is a complete listener for a parse tree produced by CpuParser.
type BaseCpuListener struct{}

var _ CpuListener = &BaseCpuListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCpuListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCpuListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCpuListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCpuListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCpuListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCpuListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCpu is called when production cpu is entered.
func (s *BaseCpuListener) EnterCpu(ctx *CpuContext) {}

// ExitCpu is called when production cpu is exited.
func (s *BaseCpuListener) ExitCpu(ctx *CpuContext) {}
