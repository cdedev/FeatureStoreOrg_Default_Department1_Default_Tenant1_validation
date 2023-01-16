// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673868556010/Carat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Carat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCaratListener is a complete listener for a parse tree produced by CaratParser.
type BaseCaratListener struct{}

var _ CaratListener = &BaseCaratListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCaratListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCaratListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCaratListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCaratListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCaratListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCaratListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCarat is called when production carat is entered.
func (s *BaseCaratListener) EnterCarat(ctx *CaratContext) {}

// ExitCarat is called when production carat is exited.
func (s *BaseCaratListener) ExitCarat(ctx *CaratContext) {}
