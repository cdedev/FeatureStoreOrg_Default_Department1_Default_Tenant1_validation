// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669709078194/Currency.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Currency

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCurrencyListener is a complete listener for a parse tree produced by CurrencyParser.
type BaseCurrencyListener struct{}

var _ CurrencyListener = &BaseCurrencyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCurrencyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCurrencyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCurrencyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCurrencyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCurrencyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCurrencyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCurrency is called when production currency is entered.
func (s *BaseCurrencyListener) EnterCurrency(ctx *CurrencyContext) {}

// ExitCurrency is called when production currency is exited.
func (s *BaseCurrencyListener) ExitCurrency(ctx *CurrencyContext) {}
