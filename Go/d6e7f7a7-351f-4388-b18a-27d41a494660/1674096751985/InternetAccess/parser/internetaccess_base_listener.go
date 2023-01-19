// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096751985/InternetAccess.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InternetAccess

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInternetAccessListener is a complete listener for a parse tree produced by InternetAccessParser.
type BaseInternetAccessListener struct{}

var _ InternetAccessListener = &BaseInternetAccessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInternetAccessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInternetAccessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInternetAccessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInternetAccessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInternetAccessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInternetAccessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInternetaccess is called when production internetaccess is entered.
func (s *BaseInternetAccessListener) EnterInternetaccess(ctx *InternetaccessContext) {}

// ExitInternetaccess is called when production internetaccess is exited.
func (s *BaseInternetAccessListener) ExitInternetaccess(ctx *InternetaccessContext) {}
