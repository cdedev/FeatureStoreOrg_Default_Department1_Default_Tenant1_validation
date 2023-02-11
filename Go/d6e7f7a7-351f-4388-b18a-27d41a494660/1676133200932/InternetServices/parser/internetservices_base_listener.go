// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676133200932/InternetServices.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // InternetServices

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInternetServicesListener is a complete listener for a parse tree produced by InternetServicesParser.
type BaseInternetServicesListener struct{}

var _ InternetServicesListener = &BaseInternetServicesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInternetServicesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInternetServicesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInternetServicesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInternetServicesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInternetServicesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInternetServicesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInternetservices is called when production internetservices is entered.
func (s *BaseInternetServicesListener) EnterInternetservices(ctx *InternetservicesContext) {}

// ExitInternetservices is called when production internetservices is exited.
func (s *BaseInternetServicesListener) ExitInternetservices(ctx *InternetservicesContext) {}
