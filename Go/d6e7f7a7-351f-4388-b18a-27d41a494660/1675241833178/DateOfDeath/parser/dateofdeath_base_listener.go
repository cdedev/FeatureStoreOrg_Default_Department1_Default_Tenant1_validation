// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675241833178/DateOfDeath.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DateOfDeath

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDateOfDeathListener is a complete listener for a parse tree produced by DateOfDeathParser.
type BaseDateOfDeathListener struct{}

var _ DateOfDeathListener = &BaseDateOfDeathListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDateOfDeathListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDateOfDeathListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDateOfDeathListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDateOfDeathListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDateOfDeathListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDateOfDeathListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDateofdeath is called when production dateofdeath is entered.
func (s *BaseDateOfDeathListener) EnterDateofdeath(ctx *DateofdeathContext) {}

// ExitDateofdeath is called when production dateofdeath is exited.
func (s *BaseDateOfDeathListener) ExitDateofdeath(ctx *DateofdeathContext) {}
