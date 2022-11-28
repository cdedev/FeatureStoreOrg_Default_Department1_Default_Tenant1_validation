// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669652405435/Failure.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Failure

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFailureListener is a complete listener for a parse tree produced by FailureParser.
type BaseFailureListener struct{}

var _ FailureListener = &BaseFailureListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFailureListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFailureListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFailureListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFailureListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFailureListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFailureListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFailure is called when production failure is entered.
func (s *BaseFailureListener) EnterFailure(ctx *FailureContext) {}

// ExitFailure is called when production failure is exited.
func (s *BaseFailureListener) ExitFailure(ctx *FailureContext) {}
