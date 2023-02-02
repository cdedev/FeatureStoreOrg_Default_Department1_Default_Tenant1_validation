// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675307081052/Systolic.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Systolic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSystolicListener is a complete listener for a parse tree produced by SystolicParser.
type BaseSystolicListener struct{}

var _ SystolicListener = &BaseSystolicListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSystolicListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSystolicListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSystolicListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSystolicListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSystolicListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSystolicListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSystolic is called when production systolic is entered.
func (s *BaseSystolicListener) EnterSystolic(ctx *SystolicContext) {}

// ExitSystolic is called when production systolic is exited.
func (s *BaseSystolicListener) ExitSystolic(ctx *SystolicContext) {}
