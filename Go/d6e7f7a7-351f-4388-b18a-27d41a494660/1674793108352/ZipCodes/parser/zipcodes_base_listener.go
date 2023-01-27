// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674793108352/ZipCodes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ZipCodes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseZipCodesListener is a complete listener for a parse tree produced by ZipCodesParser.
type BaseZipCodesListener struct{}

var _ ZipCodesListener = &BaseZipCodesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZipCodesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZipCodesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZipCodesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZipCodesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseZipCodesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseZipCodesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterZipcodes is called when production zipcodes is entered.
func (s *BaseZipCodesListener) EnterZipcodes(ctx *ZipcodesContext) {}

// ExitZipcodes is called when production zipcodes is exited.
func (s *BaseZipCodesListener) ExitZipcodes(ctx *ZipcodesContext) {}
