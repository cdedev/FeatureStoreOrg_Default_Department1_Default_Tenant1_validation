// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675836326368/Poverty.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Poverty

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePovertyListener is a complete listener for a parse tree produced by PovertyParser.
type BasePovertyListener struct{}

var _ PovertyListener = &BasePovertyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePovertyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePovertyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePovertyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePovertyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePovertyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePovertyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPoverty is called when production poverty is entered.
func (s *BasePovertyListener) EnterPoverty(ctx *PovertyContext) {}

// ExitPoverty is called when production poverty is exited.
func (s *BasePovertyListener) ExitPoverty(ctx *PovertyContext) {}
