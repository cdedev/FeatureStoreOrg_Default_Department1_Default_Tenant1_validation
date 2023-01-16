// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673867803590/Metal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Metal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMetalListener is a complete listener for a parse tree produced by MetalParser.
type BaseMetalListener struct{}

var _ MetalListener = &BaseMetalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMetalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMetalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMetalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMetalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMetalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMetalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMetal is called when production metal is entered.
func (s *BaseMetalListener) EnterMetal(ctx *MetalContext) {}

// ExitMetal is called when production metal is exited.
func (s *BaseMetalListener) ExitMetal(ctx *MetalContext) {}
