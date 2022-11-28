// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623152379/Distance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Distance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDistanceListener is a complete listener for a parse tree produced by DistanceParser.
type BaseDistanceListener struct{}

var _ DistanceListener = &BaseDistanceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDistanceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDistanceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDistanceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDistanceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDistanceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDistanceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDistance is called when production distance is entered.
func (s *BaseDistanceListener) EnterDistance(ctx *DistanceContext) {}

// ExitDistance is called when production distance is exited.
func (s *BaseDistanceListener) ExitDistance(ctx *DistanceContext) {}
