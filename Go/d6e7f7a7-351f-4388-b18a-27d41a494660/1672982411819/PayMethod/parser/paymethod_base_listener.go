// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672982411819/PayMethod.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // PayMethod

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePayMethodListener is a complete listener for a parse tree produced by PayMethodParser.
type BasePayMethodListener struct{}

var _ PayMethodListener = &BasePayMethodListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePayMethodListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePayMethodListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePayMethodListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePayMethodListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePayMethodListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePayMethodListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPaymethod is called when production paymethod is entered.
func (s *BasePayMethodListener) EnterPaymethod(ctx *PaymethodContext) {}

// ExitPaymethod is called when production paymethod is exited.
func (s *BasePayMethodListener) ExitPaymethod(ctx *PaymethodContext) {}
