// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674804602839/Depth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Depth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDepthListener is a complete listener for a parse tree produced by DepthParser.
type BaseDepthListener struct{}

var _ DepthListener = &BaseDepthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDepthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDepthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDepthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDepthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDepthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDepthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDepth is called when production depth is entered.
func (s *BaseDepthListener) EnterDepth(ctx *DepthContext) {}

// ExitDepth is called when production depth is exited.
func (s *BaseDepthListener) ExitDepth(ctx *DepthContext) {}
