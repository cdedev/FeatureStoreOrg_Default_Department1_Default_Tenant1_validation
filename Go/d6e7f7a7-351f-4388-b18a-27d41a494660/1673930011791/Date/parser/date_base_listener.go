// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673930011791/Date.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Date

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDateListener is a complete listener for a parse tree produced by DateParser.
type BaseDateListener struct{}

var _ DateListener = &BaseDateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDateListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDateListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDate is called when production date is entered.
func (s *BaseDateListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseDateListener) ExitDate(ctx *DateContext) {}
