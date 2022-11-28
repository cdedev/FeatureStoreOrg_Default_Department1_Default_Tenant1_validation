// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669656299310/Stress.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Stress

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseStressListener is a complete listener for a parse tree produced by StressParser.
type BaseStressListener struct{}

var _ StressListener = &BaseStressListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseStressListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseStressListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseStressListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseStressListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseStressListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseStressListener) ExitExpression(ctx *ExpressionContext) {}

// EnterStress is called when production stress is entered.
func (s *BaseStressListener) EnterStress(ctx *StressContext) {}

// ExitStress is called when production stress is exited.
func (s *BaseStressListener) ExitStress(ctx *StressContext) {}
