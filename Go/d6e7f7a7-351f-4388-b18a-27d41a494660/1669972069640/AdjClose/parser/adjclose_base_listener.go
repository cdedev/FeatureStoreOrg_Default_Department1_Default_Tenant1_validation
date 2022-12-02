// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669972069640/AdjClose.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AdjClose

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAdjCloseListener is a complete listener for a parse tree produced by AdjCloseParser.
type BaseAdjCloseListener struct{}

var _ AdjCloseListener = &BaseAdjCloseListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAdjCloseListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAdjCloseListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAdjCloseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAdjCloseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAdjCloseListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAdjCloseListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAdjclose is called when production adjclose is entered.
func (s *BaseAdjCloseListener) EnterAdjclose(ctx *AdjcloseContext) {}

// ExitAdjclose is called when production adjclose is exited.
func (s *BaseAdjCloseListener) ExitAdjclose(ctx *AdjcloseContext) {}
