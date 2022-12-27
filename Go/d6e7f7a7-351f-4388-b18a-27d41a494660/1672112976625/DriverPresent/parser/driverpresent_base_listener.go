// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672112976625/DriverPresent.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DriverPresent

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDriverPresentListener is a complete listener for a parse tree produced by DriverPresentParser.
type BaseDriverPresentListener struct{}

var _ DriverPresentListener = &BaseDriverPresentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDriverPresentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDriverPresentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDriverPresentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDriverPresentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDriverPresentListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDriverPresentListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDriverpresent is called when production driverpresent is entered.
func (s *BaseDriverPresentListener) EnterDriverpresent(ctx *DriverpresentContext) {}

// ExitDriverpresent is called when production driverpresent is exited.
func (s *BaseDriverPresentListener) ExitDriverpresent(ctx *DriverpresentContext) {}
