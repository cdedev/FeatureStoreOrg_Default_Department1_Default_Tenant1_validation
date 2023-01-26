// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674722475152/Occurrence.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Occurrence

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOccurrenceListener is a complete listener for a parse tree produced by OccurrenceParser.
type BaseOccurrenceListener struct{}

var _ OccurrenceListener = &BaseOccurrenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOccurrenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOccurrenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOccurrenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOccurrenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOccurrenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOccurrenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOccurrence is called when production occurrence is entered.
func (s *BaseOccurrenceListener) EnterOccurrence(ctx *OccurrenceContext) {}

// ExitOccurrence is called when production occurrence is exited.
func (s *BaseOccurrenceListener) ExitOccurrence(ctx *OccurrenceContext) {}
