// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675246511321/Offence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Offence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOffenceListener is a complete listener for a parse tree produced by OffenceParser.
type BaseOffenceListener struct{}

var _ OffenceListener = &BaseOffenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOffenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOffenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOffenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOffenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOffenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOffenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOffence is called when production offence is entered.
func (s *BaseOffenceListener) EnterOffence(ctx *OffenceContext) {}

// ExitOffence is called when production offence is exited.
func (s *BaseOffenceListener) ExitOffence(ctx *OffenceContext) {}
