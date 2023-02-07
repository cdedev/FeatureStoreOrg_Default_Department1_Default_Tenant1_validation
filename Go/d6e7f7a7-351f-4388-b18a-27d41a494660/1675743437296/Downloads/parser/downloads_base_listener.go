// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675743437296/Downloads.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Downloads

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDownloadsListener is a complete listener for a parse tree produced by DownloadsParser.
type BaseDownloadsListener struct{}

var _ DownloadsListener = &BaseDownloadsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDownloadsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDownloadsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDownloadsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDownloadsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDownloadsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDownloadsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDownloads is called when production downloads is entered.
func (s *BaseDownloadsListener) EnterDownloads(ctx *DownloadsContext) {}

// ExitDownloads is called when production downloads is exited.
func (s *BaseDownloadsListener) ExitDownloads(ctx *DownloadsContext) {}
