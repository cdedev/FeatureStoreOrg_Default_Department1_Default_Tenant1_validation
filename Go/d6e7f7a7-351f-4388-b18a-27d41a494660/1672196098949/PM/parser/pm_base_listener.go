// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672196098949/PM.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PM

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePMListener is a complete listener for a parse tree produced by PMParser.
type BasePMListener struct{}

var _ PMListener = &BasePMListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePMListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePMListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePMListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePMListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePMListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePMListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPm is called when production pm is entered.
func (s *BasePMListener) EnterPm(ctx *PmContext) {}

// ExitPm is called when production pm is exited.
func (s *BasePMListener) ExitPm(ctx *PmContext) {}
