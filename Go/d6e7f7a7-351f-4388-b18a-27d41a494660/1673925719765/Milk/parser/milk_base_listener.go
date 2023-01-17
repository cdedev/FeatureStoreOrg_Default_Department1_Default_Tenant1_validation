// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925719765/Milk.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Milk

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMilkListener is a complete listener for a parse tree produced by MilkParser.
type BaseMilkListener struct{}

var _ MilkListener = &BaseMilkListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMilkListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMilkListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMilkListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMilkListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMilkListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMilkListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMilk is called when production milk is entered.
func (s *BaseMilkListener) EnterMilk(ctx *MilkContext) {}

// ExitMilk is called when production milk is exited.
func (s *BaseMilkListener) ExitMilk(ctx *MilkContext) {}
