// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671607995826/Power.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Power

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePowerListener is a complete listener for a parse tree produced by PowerParser.
type BasePowerListener struct{}

var _ PowerListener = &BasePowerListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePowerListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePowerListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePowerListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePowerListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePowerListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePowerListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPower is called when production power is entered.
func (s *BasePowerListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BasePowerListener) ExitPower(ctx *PowerContext) {}
