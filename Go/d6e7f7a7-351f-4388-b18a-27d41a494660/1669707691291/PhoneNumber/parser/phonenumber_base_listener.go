// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669707691291/PhoneNumber.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PhoneNumber

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePhoneNumberListener is a complete listener for a parse tree produced by PhoneNumberParser.
type BasePhoneNumberListener struct{}

var _ PhoneNumberListener = &BasePhoneNumberListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePhoneNumberListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePhoneNumberListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePhoneNumberListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePhoneNumberListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePhoneNumberListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePhoneNumberListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPhonenumber is called when production phonenumber is entered.
func (s *BasePhoneNumberListener) EnterPhonenumber(ctx *PhonenumberContext) {}

// ExitPhonenumber is called when production phonenumber is exited.
func (s *BasePhoneNumberListener) ExitPhonenumber(ctx *PhonenumberContext) {}
