// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669625136425/Win.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Win

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseWinListener is a complete listener for a parse tree produced by WinParser.
type BaseWinListener struct{}

var _ WinListener = &BaseWinListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWinListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWinListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWinListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWinListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseWinListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseWinListener) ExitExpression(ctx *ExpressionContext) {}

// EnterField0 is called when production field0 is entered.
func (s *BaseWinListener) EnterField0(ctx *Field0Context) {}

// ExitField0 is called when production field0 is exited.
func (s *BaseWinListener) ExitField0(ctx *Field0Context) {}
