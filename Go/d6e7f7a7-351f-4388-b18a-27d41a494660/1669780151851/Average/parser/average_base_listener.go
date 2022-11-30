// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669780151851/Average.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Average

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAverageListener is a complete listener for a parse tree produced by AverageParser.
type BaseAverageListener struct{}

var _ AverageListener = &BaseAverageListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAverageListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAverageListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAverageListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAverageListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAverageListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAverageListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAverage is called when production average is entered.
func (s *BaseAverageListener) EnterAverage(ctx *AverageContext) {}

// ExitAverage is called when production average is exited.
func (s *BaseAverageListener) ExitAverage(ctx *AverageContext) {}
