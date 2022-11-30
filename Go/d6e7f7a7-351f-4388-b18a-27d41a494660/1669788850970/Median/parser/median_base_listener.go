// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669788850970/Median.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Median

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMedianListener is a complete listener for a parse tree produced by MedianParser.
type BaseMedianListener struct{}

var _ MedianListener = &BaseMedianListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMedianListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMedianListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMedianListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMedianListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMedianListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMedianListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMedian is called when production median is entered.
func (s *BaseMedianListener) EnterMedian(ctx *MedianContext) {}

// ExitMedian is called when production median is exited.
func (s *BaseMedianListener) ExitMedian(ctx *MedianContext) {}
