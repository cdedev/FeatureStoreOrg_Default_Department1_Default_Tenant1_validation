// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673237722630/Inhabitant.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Inhabitant

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInhabitantListener is a complete listener for a parse tree produced by InhabitantParser.
type BaseInhabitantListener struct{}

var _ InhabitantListener = &BaseInhabitantListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInhabitantListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInhabitantListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInhabitantListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInhabitantListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInhabitantListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInhabitantListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInhabitant is called when production inhabitant is entered.
func (s *BaseInhabitantListener) EnterInhabitant(ctx *InhabitantContext) {}

// ExitInhabitant is called when production inhabitant is exited.
func (s *BaseInhabitantListener) ExitInhabitant(ctx *InhabitantContext) {}
