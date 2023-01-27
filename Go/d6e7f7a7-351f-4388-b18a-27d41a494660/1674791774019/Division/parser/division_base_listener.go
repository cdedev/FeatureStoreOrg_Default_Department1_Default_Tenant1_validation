// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674791774019/Division.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Division

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDivisionListener is a complete listener for a parse tree produced by DivisionParser.
type BaseDivisionListener struct{}

var _ DivisionListener = &BaseDivisionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDivisionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDivisionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDivisionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDivisionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDivisionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDivisionListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDivision is called when production division is entered.
func (s *BaseDivisionListener) EnterDivision(ctx *DivisionContext) {}

// ExitDivision is called when production division is exited.
func (s *BaseDivisionListener) ExitDivision(ctx *DivisionContext) {}
