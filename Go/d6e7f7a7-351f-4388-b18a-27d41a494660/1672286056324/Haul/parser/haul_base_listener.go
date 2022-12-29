// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672286056324/Haul.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Haul

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHaulListener is a complete listener for a parse tree produced by HaulParser.
type BaseHaulListener struct{}

var _ HaulListener = &BaseHaulListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHaulListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHaulListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHaulListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHaulListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHaulListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHaulListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHaul is called when production haul is entered.
func (s *BaseHaulListener) EnterHaul(ctx *HaulContext) {}

// ExitHaul is called when production haul is exited.
func (s *BaseHaulListener) ExitHaul(ctx *HaulContext) {}
