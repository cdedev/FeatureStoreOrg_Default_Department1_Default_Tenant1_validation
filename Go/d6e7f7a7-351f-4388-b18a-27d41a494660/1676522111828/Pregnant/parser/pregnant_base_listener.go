// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676522111828/Pregnant.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pregnant

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePregnantListener is a complete listener for a parse tree produced by PregnantParser.
type BasePregnantListener struct{}

var _ PregnantListener = &BasePregnantListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePregnantListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePregnantListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePregnantListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePregnantListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePregnantListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePregnantListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPregnant is called when production pregnant is entered.
func (s *BasePregnantListener) EnterPregnant(ctx *PregnantContext) {}

// ExitPregnant is called when production pregnant is exited.
func (s *BasePregnantListener) ExitPregnant(ctx *PregnantContext) {}
