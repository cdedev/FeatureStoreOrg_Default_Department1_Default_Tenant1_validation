// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673676416291/Smoothness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Smoothness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSmoothnessListener is a complete listener for a parse tree produced by SmoothnessParser.
type BaseSmoothnessListener struct{}

var _ SmoothnessListener = &BaseSmoothnessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSmoothnessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSmoothnessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSmoothnessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSmoothnessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSmoothnessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSmoothnessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSmoothness is called when production smoothness is entered.
func (s *BaseSmoothnessListener) EnterSmoothness(ctx *SmoothnessContext) {}

// ExitSmoothness is called when production smoothness is exited.
func (s *BaseSmoothnessListener) ExitSmoothness(ctx *SmoothnessContext) {}
