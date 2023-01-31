// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675149663663/Route.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Route

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRouteListener is a complete listener for a parse tree produced by RouteParser.
type BaseRouteListener struct{}

var _ RouteListener = &BaseRouteListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRouteListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRouteListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRouteListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRouteListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRouteListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRouteListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRoute is called when production route is entered.
func (s *BaseRouteListener) EnterRoute(ctx *RouteContext) {}

// ExitRoute is called when production route is exited.
func (s *BaseRouteListener) ExitRoute(ctx *RouteContext) {}
