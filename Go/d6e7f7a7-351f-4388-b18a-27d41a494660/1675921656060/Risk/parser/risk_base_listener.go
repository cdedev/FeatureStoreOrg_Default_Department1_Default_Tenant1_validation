// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675921656060/Risk.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Risk

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRiskListener is a complete listener for a parse tree produced by RiskParser.
type BaseRiskListener struct{}

var _ RiskListener = &BaseRiskListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRiskListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRiskListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRiskListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRiskListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRiskListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRiskListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRisk is called when production risk is entered.
func (s *BaseRiskListener) EnterRisk(ctx *RiskContext) {}

// ExitRisk is called when production risk is exited.
func (s *BaseRiskListener) ExitRisk(ctx *RiskContext) {}
