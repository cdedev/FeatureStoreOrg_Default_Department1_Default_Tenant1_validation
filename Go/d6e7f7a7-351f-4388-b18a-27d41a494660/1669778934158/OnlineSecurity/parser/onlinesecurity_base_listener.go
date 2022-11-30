// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778934158/OnlineSecurity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OnlineSecurity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOnlineSecurityListener is a complete listener for a parse tree produced by OnlineSecurityParser.
type BaseOnlineSecurityListener struct{}

var _ OnlineSecurityListener = &BaseOnlineSecurityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOnlineSecurityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOnlineSecurityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOnlineSecurityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOnlineSecurityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOnlineSecurityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOnlineSecurityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOnlinesecurity is called when production onlinesecurity is entered.
func (s *BaseOnlineSecurityListener) EnterOnlinesecurity(ctx *OnlinesecurityContext) {}

// ExitOnlinesecurity is called when production onlinesecurity is exited.
func (s *BaseOnlineSecurityListener) ExitOnlinesecurity(ctx *OnlinesecurityContext) {}
