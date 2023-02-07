// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675744972021/Founder.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Founder

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFounderListener is a complete listener for a parse tree produced by FounderParser.
type BaseFounderListener struct{}

var _ FounderListener = &BaseFounderListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFounderListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFounderListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFounderListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFounderListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFounderListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFounderListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFounder is called when production founder is entered.
func (s *BaseFounderListener) EnterFounder(ctx *FounderContext) {}

// ExitFounder is called when production founder is exited.
func (s *BaseFounderListener) ExitFounder(ctx *FounderContext) {}
