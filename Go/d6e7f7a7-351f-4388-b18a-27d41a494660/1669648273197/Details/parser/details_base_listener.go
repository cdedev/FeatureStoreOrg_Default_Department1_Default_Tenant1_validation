// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669648273197/Details.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Details

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDetailsListener is a complete listener for a parse tree produced by DetailsParser.
type BaseDetailsListener struct{}

var _ DetailsListener = &BaseDetailsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDetailsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDetailsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDetailsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDetailsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDetailsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDetailsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDetails is called when production details is entered.
func (s *BaseDetailsListener) EnterDetails(ctx *DetailsContext) {}

// ExitDetails is called when production details is exited.
func (s *BaseDetailsListener) ExitDetails(ctx *DetailsContext) {}
