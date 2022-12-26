// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672075300568/CanopyClass.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CanopyClass

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCanopyClassListener is a complete listener for a parse tree produced by CanopyClassParser.
type BaseCanopyClassListener struct{}

var _ CanopyClassListener = &BaseCanopyClassListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCanopyClassListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCanopyClassListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCanopyClassListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCanopyClassListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCanopyClassListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCanopyClassListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCanopyclass is called when production canopyclass is entered.
func (s *BaseCanopyClassListener) EnterCanopyclass(ctx *CanopyclassContext) {}

// ExitCanopyclass is called when production canopyclass is exited.
func (s *BaseCanopyClassListener) ExitCanopyclass(ctx *CanopyclassContext) {}
