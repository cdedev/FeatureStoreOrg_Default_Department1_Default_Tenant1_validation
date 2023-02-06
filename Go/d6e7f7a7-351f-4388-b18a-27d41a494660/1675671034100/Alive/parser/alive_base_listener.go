// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675671034100/Alive.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Alive

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAliveListener is a complete listener for a parse tree produced by AliveParser.
type BaseAliveListener struct{}

var _ AliveListener = &BaseAliveListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAliveListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAliveListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAliveListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAliveListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAliveListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAliveListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAlive is called when production alive is entered.
func (s *BaseAliveListener) EnterAlive(ctx *AliveContext) {}

// ExitAlive is called when production alive is exited.
func (s *BaseAliveListener) ExitAlive(ctx *AliveContext) {}
