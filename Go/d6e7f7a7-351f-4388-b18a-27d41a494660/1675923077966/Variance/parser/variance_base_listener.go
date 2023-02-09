// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675923077966/Variance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Variance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVarianceListener is a complete listener for a parse tree produced by VarianceParser.
type BaseVarianceListener struct{}

var _ VarianceListener = &BaseVarianceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVarianceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVarianceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVarianceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVarianceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVarianceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVarianceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVariance is called when production variance is entered.
func (s *BaseVarianceListener) EnterVariance(ctx *VarianceContext) {}

// ExitVariance is called when production variance is exited.
func (s *BaseVarianceListener) ExitVariance(ctx *VarianceContext) {}
