// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673237785511/Iron.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Iron

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseIronListener is a complete listener for a parse tree produced by IronParser.
type BaseIronListener struct{}

var _ IronListener = &BaseIronListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseIronListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseIronListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseIronListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseIronListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseIronListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseIronListener) ExitExpression(ctx *ExpressionContext) {}

// EnterIron is called when production iron is entered.
func (s *BaseIronListener) EnterIron(ctx *IronContext) {}

// ExitIron is called when production iron is exited.
func (s *BaseIronListener) ExitIron(ctx *IronContext) {}
