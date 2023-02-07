// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675740638061/DollarValue.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DollarValue

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDollarValueListener is a complete listener for a parse tree produced by DollarValueParser.
type BaseDollarValueListener struct{}

var _ DollarValueListener = &BaseDollarValueListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDollarValueListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDollarValueListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDollarValueListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDollarValueListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDollarValueListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDollarValueListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDollarvalue is called when production dollarvalue is entered.
func (s *BaseDollarValueListener) EnterDollarvalue(ctx *DollarvalueContext) {}

// ExitDollarvalue is called when production dollarvalue is exited.
func (s *BaseDollarValueListener) ExitDollarvalue(ctx *DollarvalueContext) {}
