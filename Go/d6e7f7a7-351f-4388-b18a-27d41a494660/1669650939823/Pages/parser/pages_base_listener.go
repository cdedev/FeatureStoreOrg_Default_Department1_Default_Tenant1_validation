// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669650939823/Pages.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pages

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePagesListener is a complete listener for a parse tree produced by PagesParser.
type BasePagesListener struct{}

var _ PagesListener = &BasePagesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePagesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePagesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePagesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePagesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePagesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePagesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPages is called when production pages is entered.
func (s *BasePagesListener) EnterPages(ctx *PagesContext) {}

// ExitPages is called when production pages is exited.
func (s *BasePagesListener) ExitPages(ctx *PagesContext) {}
