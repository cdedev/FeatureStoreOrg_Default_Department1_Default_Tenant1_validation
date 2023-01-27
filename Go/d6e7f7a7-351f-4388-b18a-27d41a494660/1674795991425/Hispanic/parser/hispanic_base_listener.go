// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674795991425/Hispanic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hispanic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHispanicListener is a complete listener for a parse tree produced by HispanicParser.
type BaseHispanicListener struct{}

var _ HispanicListener = &BaseHispanicListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHispanicListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHispanicListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHispanicListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHispanicListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHispanicListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHispanicListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHispanic is called when production hispanic is entered.
func (s *BaseHispanicListener) EnterHispanic(ctx *HispanicContext) {}

// ExitHispanic is called when production hispanic is exited.
func (s *BaseHispanicListener) ExitHispanic(ctx *HispanicContext) {}
