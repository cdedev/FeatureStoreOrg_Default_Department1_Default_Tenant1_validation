// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671608397452/Ram.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ram

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRamListener is a complete listener for a parse tree produced by RamParser.
type BaseRamListener struct{}

var _ RamListener = &BaseRamListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRamListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRamListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRamListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRamListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRamListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRamListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRam is called when production ram is entered.
func (s *BaseRamListener) EnterRam(ctx *RamContext) {}

// ExitRam is called when production ram is exited.
func (s *BaseRamListener) ExitRam(ctx *RamContext) {}
