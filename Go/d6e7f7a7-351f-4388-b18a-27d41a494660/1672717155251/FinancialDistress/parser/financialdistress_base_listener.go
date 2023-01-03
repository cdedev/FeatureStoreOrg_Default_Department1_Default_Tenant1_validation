// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672717155251/FinancialDistress.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FinancialDistress

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFinancialDistressListener is a complete listener for a parse tree produced by FinancialDistressParser.
type BaseFinancialDistressListener struct{}

var _ FinancialDistressListener = &BaseFinancialDistressListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFinancialDistressListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFinancialDistressListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFinancialDistressListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFinancialDistressListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFinancialDistressListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFinancialDistressListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFinancialdistress is called when production financialdistress is entered.
func (s *BaseFinancialDistressListener) EnterFinancialdistress(ctx *FinancialdistressContext) {}

// ExitFinancialdistress is called when production financialdistress is exited.
func (s *BaseFinancialDistressListener) ExitFinancialdistress(ctx *FinancialdistressContext) {}
