// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671207762977/ConfirmedCases.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ConfirmedCases

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConfirmedCasesListener is a complete listener for a parse tree produced by ConfirmedCasesParser.
type BaseConfirmedCasesListener struct{}

var _ ConfirmedCasesListener = &BaseConfirmedCasesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConfirmedCasesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConfirmedCasesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConfirmedCasesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConfirmedCasesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConfirmedCasesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConfirmedCasesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConfirmedcases is called when production confirmedcases is entered.
func (s *BaseConfirmedCasesListener) EnterConfirmedcases(ctx *ConfirmedcasesContext) {}

// ExitConfirmedcases is called when production confirmedcases is exited.
func (s *BaseConfirmedCasesListener) ExitConfirmedcases(ctx *ConfirmedcasesContext) {}
