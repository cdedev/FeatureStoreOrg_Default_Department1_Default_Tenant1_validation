// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675703060989/Points.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Points

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePointsListener is a complete listener for a parse tree produced by PointsParser.
type BasePointsListener struct{}

var _ PointsListener = &BasePointsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePointsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePointsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePointsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePointsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePointsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePointsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPoints is called when production points is entered.
func (s *BasePointsListener) EnterPoints(ctx *PointsContext) {}

// ExitPoints is called when production points is exited.
func (s *BasePointsListener) ExitPoints(ctx *PointsContext) {}
