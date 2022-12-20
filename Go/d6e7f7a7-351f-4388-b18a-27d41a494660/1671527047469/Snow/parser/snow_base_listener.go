// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671527047469/Snow.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Snow

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSnowListener is a complete listener for a parse tree produced by SnowParser.
type BaseSnowListener struct{}

var _ SnowListener = &BaseSnowListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSnowListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSnowListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSnowListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSnowListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSnowListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSnowListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSnow is called when production snow is entered.
func (s *BaseSnowListener) EnterSnow(ctx *SnowContext) {}

// ExitSnow is called when production snow is exited.
func (s *BaseSnowListener) ExitSnow(ctx *SnowContext) {}
