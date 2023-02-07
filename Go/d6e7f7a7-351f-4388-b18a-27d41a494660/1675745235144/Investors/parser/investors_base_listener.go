// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675745235144/Investors.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Investors

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInvestorsListener is a complete listener for a parse tree produced by InvestorsParser.
type BaseInvestorsListener struct{}

var _ InvestorsListener = &BaseInvestorsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInvestorsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInvestorsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInvestorsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInvestorsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInvestorsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInvestorsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInvestors is called when production investors is entered.
func (s *BaseInvestorsListener) EnterInvestors(ctx *InvestorsContext) {}

// ExitInvestors is called when production investors is exited.
func (s *BaseInvestorsListener) ExitInvestors(ctx *InvestorsContext) {}
