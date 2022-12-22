// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671690140451/Occupation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Occupation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOccupationListener is a complete listener for a parse tree produced by OccupationParser.
type BaseOccupationListener struct{}

var _ OccupationListener = &BaseOccupationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOccupationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOccupationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOccupationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOccupationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOccupationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOccupationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOccupation is called when production occupation is entered.
func (s *BaseOccupationListener) EnterOccupation(ctx *OccupationContext) {}

// ExitOccupation is called when production occupation is exited.
func (s *BaseOccupationListener) ExitOccupation(ctx *OccupationContext) {}
