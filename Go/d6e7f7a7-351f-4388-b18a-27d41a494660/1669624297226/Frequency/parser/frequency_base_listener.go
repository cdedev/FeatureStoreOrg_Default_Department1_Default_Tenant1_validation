// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669624297226/Frequency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Frequency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFrequencyListener is a complete listener for a parse tree produced by FrequencyParser.
type BaseFrequencyListener struct{}

var _ FrequencyListener = &BaseFrequencyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFrequencyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFrequencyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFrequencyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFrequencyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFrequencyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFrequencyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFrequency is called when production frequency is entered.
func (s *BaseFrequencyListener) EnterFrequency(ctx *FrequencyContext) {}

// ExitFrequency is called when production frequency is exited.
func (s *BaseFrequencyListener) ExitFrequency(ctx *FrequencyContext) {}
