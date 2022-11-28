// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669654144152/Worklifebalance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Worklifebalance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWorklifebalanceListener is a complete listener for a parse tree produced by WorklifebalanceParser.
type BaseWorklifebalanceListener struct{}

var _ WorklifebalanceListener = &BaseWorklifebalanceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWorklifebalanceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWorklifebalanceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWorklifebalanceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWorklifebalanceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWorklifebalanceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWorklifebalanceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterWorklifebalance is called when production worklifebalance is entered.
func (s *BaseWorklifebalanceListener) EnterWorklifebalance(ctx *WorklifebalanceContext) {}

// ExitWorklifebalance is called when production worklifebalance is exited.
func (s *BaseWorklifebalanceListener) ExitWorklifebalance(ctx *WorklifebalanceContext) {}
