// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674793042480/Wards.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Wards

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWardsListener is a complete listener for a parse tree produced by WardsParser.
type BaseWardsListener struct{}

var _ WardsListener = &BaseWardsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWardsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWardsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWardsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWardsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWardsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWardsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWards is called when production wards is entered.
func (s *BaseWardsListener) EnterWards(ctx *WardsContext) {}

// ExitWards is called when production wards is exited.
func (s *BaseWardsListener) ExitWards(ctx *WardsContext) {}
