// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673928456995/Area.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Area

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAreaListener is a complete listener for a parse tree produced by AreaParser.
type BaseAreaListener struct{}

var _ AreaListener = &BaseAreaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAreaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAreaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAreaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAreaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAreaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAreaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterArea is called when production area is entered.
func (s *BaseAreaListener) EnterArea(ctx *AreaContext) {}

// ExitArea is called when production area is exited.
func (s *BaseAreaListener) ExitArea(ctx *AreaContext) {}
