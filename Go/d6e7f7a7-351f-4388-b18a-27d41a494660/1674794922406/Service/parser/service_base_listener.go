// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674794922406/Service.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Service

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseServiceListener is a complete listener for a parse tree produced by ServiceParser.
type BaseServiceListener struct{}

var _ ServiceListener = &BaseServiceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseServiceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseServiceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseServiceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseServiceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseServiceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseServiceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterService is called when production service is entered.
func (s *BaseServiceListener) EnterService(ctx *ServiceContext) {}

// ExitService is called when production service is exited.
func (s *BaseServiceListener) ExitService(ctx *ServiceContext) {}
