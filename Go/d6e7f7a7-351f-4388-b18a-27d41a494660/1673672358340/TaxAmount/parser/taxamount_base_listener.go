// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673672358340/TaxAmount.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TaxAmount

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTaxAmountListener is a complete listener for a parse tree produced by TaxAmountParser.
type BaseTaxAmountListener struct{}

var _ TaxAmountListener = &BaseTaxAmountListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTaxAmountListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTaxAmountListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTaxAmountListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTaxAmountListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTaxAmountListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTaxAmountListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTaxamount is called when production taxamount is entered.
func (s *BaseTaxAmountListener) EnterTaxamount(ctx *TaxamountContext) {}

// ExitTaxamount is called when production taxamount is exited.
func (s *BaseTaxAmountListener) ExitTaxamount(ctx *TaxamountContext) {}
