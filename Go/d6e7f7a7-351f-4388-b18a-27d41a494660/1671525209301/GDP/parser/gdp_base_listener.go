// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671525209301/GDP.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GDP

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGDPListener is a complete listener for a parse tree produced by GDPParser.
type BaseGDPListener struct{}

var _ GDPListener = &BaseGDPListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGDPListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGDPListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGDPListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGDPListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGDPListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGDPListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGdp is called when production gdp is entered.
func (s *BaseGDPListener) EnterGdp(ctx *GdpContext) {}

// ExitGdp is called when production gdp is exited.
func (s *BaseGDPListener) ExitGdp(ctx *GdpContext) {}
