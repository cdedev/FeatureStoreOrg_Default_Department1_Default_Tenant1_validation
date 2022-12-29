// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672285317914/Detail.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Detail

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDetailListener is a complete listener for a parse tree produced by DetailParser.
type BaseDetailListener struct{}

var _ DetailListener = &BaseDetailListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDetailListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDetailListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDetailListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDetailListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDetailListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDetailListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDetail is called when production detail is entered.
func (s *BaseDetailListener) EnterDetail(ctx *DetailContext) {}

// ExitDetail is called when production detail is exited.
func (s *BaseDetailListener) ExitDetail(ctx *DetailContext) {}
