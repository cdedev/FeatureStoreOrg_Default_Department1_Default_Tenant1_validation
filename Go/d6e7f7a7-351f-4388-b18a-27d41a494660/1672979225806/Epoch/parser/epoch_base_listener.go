// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672979225806/Epoch.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Epoch

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEpochListener is a complete listener for a parse tree produced by EpochParser.
type BaseEpochListener struct{}

var _ EpochListener = &BaseEpochListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEpochListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEpochListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEpochListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEpochListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEpochListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEpochListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEpoch is called when production epoch is entered.
func (s *BaseEpochListener) EnterEpoch(ctx *EpochContext) {}

// ExitEpoch is called when production epoch is exited.
func (s *BaseEpochListener) ExitEpoch(ctx *EpochContext) {}
