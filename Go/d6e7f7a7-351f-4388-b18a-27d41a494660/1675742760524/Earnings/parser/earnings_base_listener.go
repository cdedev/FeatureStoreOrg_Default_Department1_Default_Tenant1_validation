// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675742760524/Earnings.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Earnings

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEarningsListener is a complete listener for a parse tree produced by EarningsParser.
type BaseEarningsListener struct{}

var _ EarningsListener = &BaseEarningsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEarningsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEarningsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEarningsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEarningsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEarningsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEarningsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEarnings is called when production earnings is entered.
func (s *BaseEarningsListener) EnterEarnings(ctx *EarningsContext) {}

// ExitEarnings is called when production earnings is exited.
func (s *BaseEarningsListener) ExitEarnings(ctx *EarningsContext) {}
