// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672722988491/Nationality.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nationality

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNationalityListener is a complete listener for a parse tree produced by NationalityParser.
type BaseNationalityListener struct{}

var _ NationalityListener = &BaseNationalityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNationalityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNationalityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNationalityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNationalityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNationalityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNationalityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNationality is called when production nationality is entered.
func (s *BaseNationalityListener) EnterNationality(ctx *NationalityContext) {}

// ExitNationality is called when production nationality is exited.
func (s *BaseNationalityListener) ExitNationality(ctx *NationalityContext) {}
