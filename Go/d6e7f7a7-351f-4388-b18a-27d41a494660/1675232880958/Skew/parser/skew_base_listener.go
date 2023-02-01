// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675232880958/Skew.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Skew

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSkewListener is a complete listener for a parse tree produced by SkewParser.
type BaseSkewListener struct{}

var _ SkewListener = &BaseSkewListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSkewListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSkewListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSkewListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSkewListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSkewListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSkewListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSkew is called when production skew is entered.
func (s *BaseSkewListener) EnterSkew(ctx *SkewContext) {}

// ExitSkew is called when production skew is exited.
func (s *BaseSkewListener) ExitSkew(ctx *SkewContext) {}
