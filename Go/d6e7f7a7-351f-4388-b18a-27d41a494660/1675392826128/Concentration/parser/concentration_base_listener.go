// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675392826128/Concentration.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Concentration

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConcentrationListener is a complete listener for a parse tree produced by ConcentrationParser.
type BaseConcentrationListener struct{}

var _ ConcentrationListener = &BaseConcentrationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConcentrationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConcentrationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConcentrationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConcentrationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConcentrationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConcentrationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConcentration is called when production concentration is entered.
func (s *BaseConcentrationListener) EnterConcentration(ctx *ConcentrationContext) {}

// ExitConcentration is called when production concentration is exited.
func (s *BaseConcentrationListener) ExitConcentration(ctx *ConcentrationContext) {}
