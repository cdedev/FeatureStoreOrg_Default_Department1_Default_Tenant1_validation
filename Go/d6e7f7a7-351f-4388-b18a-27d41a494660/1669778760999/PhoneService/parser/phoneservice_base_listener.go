// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778760999/PhoneService.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PhoneService

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePhoneServiceListener is a complete listener for a parse tree produced by PhoneServiceParser.
type BasePhoneServiceListener struct{}

var _ PhoneServiceListener = &BasePhoneServiceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePhoneServiceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePhoneServiceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePhoneServiceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePhoneServiceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePhoneServiceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePhoneServiceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPhoneservice is called when production phoneservice is entered.
func (s *BasePhoneServiceListener) EnterPhoneservice(ctx *PhoneserviceContext) {}

// ExitPhoneservice is called when production phoneservice is exited.
func (s *BasePhoneServiceListener) ExitPhoneservice(ctx *PhoneserviceContext) {}
