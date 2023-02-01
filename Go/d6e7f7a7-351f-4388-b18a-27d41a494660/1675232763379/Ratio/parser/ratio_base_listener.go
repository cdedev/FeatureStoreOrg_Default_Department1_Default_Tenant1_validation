// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675232763379/Ratio.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Ratio

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRatioListener is a complete listener for a parse tree produced by RatioParser.
type BaseRatioListener struct{}

var _ RatioListener = &BaseRatioListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRatioListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRatioListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRatioListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRatioListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRatioListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRatioListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRatio is called when production ratio is entered.
func (s *BaseRatioListener) EnterRatio(ctx *RatioContext) {}

// ExitRatio is called when production ratio is exited.
func (s *BaseRatioListener) ExitRatio(ctx *RatioContext) {}
