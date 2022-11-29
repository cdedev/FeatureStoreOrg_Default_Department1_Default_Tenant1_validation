// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669703288157/Price.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Price

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePriceListener is a complete listener for a parse tree produced by PriceParser.
type BasePriceListener struct{}

var _ PriceListener = &BasePriceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePriceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePriceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePriceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePriceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePriceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePriceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPrice is called when production price is entered.
func (s *BasePriceListener) EnterPrice(ctx *PriceContext) {}

// ExitPrice is called when production price is exited.
func (s *BasePriceListener) ExitPrice(ctx *PriceContext) {}
