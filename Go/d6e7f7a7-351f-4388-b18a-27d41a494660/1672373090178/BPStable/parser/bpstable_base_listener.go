// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672373090178/BPStable.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BPStable

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBPStableListener is a complete listener for a parse tree produced by BPStableParser.
type BaseBPStableListener struct{}

var _ BPStableListener = &BaseBPStableListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBPStableListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBPStableListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBPStableListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBPStableListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBPStableListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBPStableListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBpstable is called when production bpstable is entered.
func (s *BaseBPStableListener) EnterBpstable(ctx *BpstableContext) {}

// ExitBpstable is called when production bpstable is exited.
func (s *BaseBPStableListener) ExitBpstable(ctx *BpstableContext) {}
