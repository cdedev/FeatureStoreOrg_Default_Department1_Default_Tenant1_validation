// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623542984/SiteLimit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SiteLimit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSiteLimitListener is a complete listener for a parse tree produced by SiteLimitParser.
type BaseSiteLimitListener struct{}

var _ SiteLimitListener = &BaseSiteLimitListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSiteLimitListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSiteLimitListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSiteLimitListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSiteLimitListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSiteLimitListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSiteLimitListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSitelimit is called when production sitelimit is entered.
func (s *BaseSiteLimitListener) EnterSitelimit(ctx *SitelimitContext) {}

// ExitSitelimit is called when production sitelimit is exited.
func (s *BaseSiteLimitListener) ExitSitelimit(ctx *SitelimitContext) {}
