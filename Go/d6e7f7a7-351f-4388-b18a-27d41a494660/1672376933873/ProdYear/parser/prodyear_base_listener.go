// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376933873/ProdYear.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ProdYear

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseProdYearListener is a complete listener for a parse tree produced by ProdYearParser.
type BaseProdYearListener struct{}

var _ ProdYearListener = &BaseProdYearListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseProdYearListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseProdYearListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseProdYearListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseProdYearListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseProdYearListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseProdYearListener) ExitExpression(ctx *ExpressionContext) {}

// EnterProdyear is called when production prodyear is entered.
func (s *BaseProdYearListener) EnterProdyear(ctx *ProdyearContext) {}

// ExitProdyear is called when production prodyear is exited.
func (s *BaseProdYearListener) ExitProdyear(ctx *ProdyearContext) {}
