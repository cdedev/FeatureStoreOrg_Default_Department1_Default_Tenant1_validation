// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669790010302/Delta.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Delta

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDeltaListener is a complete listener for a parse tree produced by DeltaParser.
type BaseDeltaListener struct{}

var _ DeltaListener = &BaseDeltaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDeltaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDeltaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDeltaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDeltaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDeltaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDeltaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDelta is called when production delta is entered.
func (s *BaseDeltaListener) EnterDelta(ctx *DeltaContext) {}

// ExitDelta is called when production delta is exited.
func (s *BaseDeltaListener) ExitDelta(ctx *DeltaContext) {}
