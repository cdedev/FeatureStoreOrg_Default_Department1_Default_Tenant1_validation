// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522166037/Sick.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sick

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSickListener is a complete listener for a parse tree produced by SickParser.
type BaseSickListener struct{}

var _ SickListener = &BaseSickListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSickListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSickListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSickListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSickListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSickListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSickListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSick is called when production sick is entered.
func (s *BaseSickListener) EnterSick(ctx *SickContext) {}

// ExitSick is called when production sick is exited.
func (s *BaseSickListener) ExitSick(ctx *SickContext) {}
