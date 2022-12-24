// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671867326616/CurrencyCrisis.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // CurrencyCrisis

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCurrencyCrisisListener is a complete listener for a parse tree produced by CurrencyCrisisParser.
type BaseCurrencyCrisisListener struct{}

var _ CurrencyCrisisListener = &BaseCurrencyCrisisListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCurrencyCrisisListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCurrencyCrisisListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCurrencyCrisisListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCurrencyCrisisListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCurrencyCrisisListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCurrencyCrisisListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCurrencycrisis is called when production currencycrisis is entered.
func (s *BaseCurrencyCrisisListener) EnterCurrencycrisis(ctx *CurrencycrisisContext) {}

// ExitCurrencycrisis is called when production currencycrisis is exited.
func (s *BaseCurrencyCrisisListener) ExitCurrencycrisis(ctx *CurrencycrisisContext) {}
