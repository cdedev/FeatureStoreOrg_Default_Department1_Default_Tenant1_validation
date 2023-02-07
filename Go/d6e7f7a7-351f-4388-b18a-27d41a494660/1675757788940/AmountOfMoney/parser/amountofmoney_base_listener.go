// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675757788940/AmountOfMoney.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AmountOfMoney

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAmountOfMoneyListener is a complete listener for a parse tree produced by AmountOfMoneyParser.
type BaseAmountOfMoneyListener struct{}

var _ AmountOfMoneyListener = &BaseAmountOfMoneyListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAmountOfMoneyListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAmountOfMoneyListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAmountOfMoneyListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAmountOfMoneyListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAmountOfMoneyListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAmountOfMoneyListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAmountofmoney is called when production amountofmoney is entered.
func (s *BaseAmountOfMoneyListener) EnterAmountofmoney(ctx *AmountofmoneyContext) {}

// ExitAmountofmoney is called when production amountofmoney is exited.
func (s *BaseAmountOfMoneyListener) ExitAmountofmoney(ctx *AmountofmoneyContext) {}
