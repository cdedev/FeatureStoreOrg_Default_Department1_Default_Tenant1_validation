// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669707921750/ZipCode.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ZipCode

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseZipCodeListener is a complete listener for a parse tree produced by ZipCodeParser.
type BaseZipCodeListener struct{}

var _ ZipCodeListener = &BaseZipCodeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZipCodeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZipCodeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZipCodeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZipCodeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseZipCodeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseZipCodeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterZipcode is called when production zipcode is entered.
func (s *BaseZipCodeListener) EnterZipcode(ctx *ZipcodeContext) {}

// ExitZipcode is called when production zipcode is exited.
func (s *BaseZipCodeListener) ExitZipcode(ctx *ZipcodeContext) {}
