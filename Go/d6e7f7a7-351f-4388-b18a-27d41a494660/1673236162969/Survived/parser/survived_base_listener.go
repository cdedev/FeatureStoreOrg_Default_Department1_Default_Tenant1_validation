// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673236162969/Survived.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Survived

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSurvivedListener is a complete listener for a parse tree produced by SurvivedParser.
type BaseSurvivedListener struct{}

var _ SurvivedListener = &BaseSurvivedListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSurvivedListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSurvivedListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSurvivedListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSurvivedListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSurvivedListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSurvivedListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSurvived is called when production survived is entered.
func (s *BaseSurvivedListener) EnterSurvived(ctx *SurvivedContext) {}

// ExitSurvived is called when production survived is exited.
func (s *BaseSurvivedListener) ExitSurvived(ctx *SurvivedContext) {}
