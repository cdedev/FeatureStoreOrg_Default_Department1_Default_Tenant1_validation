// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288638495/Year.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Year

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseYearListener is a complete listener for a parse tree produced by YearParser.
type BaseYearListener struct{}

var _ YearListener = &BaseYearListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseYearListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseYearListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseYearListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseYearListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseYearListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseYearListener) ExitExpression(ctx *ExpressionContext) {}

// EnterYear is called when production year is entered.
func (s *BaseYearListener) EnterYear(ctx *YearContext) {}

// ExitYear is called when production year is exited.
func (s *BaseYearListener) ExitYear(ctx *YearContext) {}
