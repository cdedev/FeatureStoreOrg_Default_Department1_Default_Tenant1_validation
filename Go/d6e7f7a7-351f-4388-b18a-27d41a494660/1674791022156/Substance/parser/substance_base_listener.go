// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674791022156/Substance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Substance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSubstanceListener is a complete listener for a parse tree produced by SubstanceParser.
type BaseSubstanceListener struct{}

var _ SubstanceListener = &BaseSubstanceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSubstanceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSubstanceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSubstanceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSubstanceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSubstanceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSubstanceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSubstance is called when production substance is entered.
func (s *BaseSubstanceListener) EnterSubstance(ctx *SubstanceContext) {}

// ExitSubstance is called when production substance is exited.
func (s *BaseSubstanceListener) ExitSubstance(ctx *SubstanceContext) {}
