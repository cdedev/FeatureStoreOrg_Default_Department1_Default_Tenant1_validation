// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672638016123/TransTotal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TransTotal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTransTotalListener is a complete listener for a parse tree produced by TransTotalParser.
type BaseTransTotalListener struct{}

var _ TransTotalListener = &BaseTransTotalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTransTotalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTransTotalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTransTotalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTransTotalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTransTotalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTransTotalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTranstotal is called when production transtotal is entered.
func (s *BaseTransTotalListener) EnterTranstotal(ctx *TranstotalContext) {}

// ExitTranstotal is called when production transtotal is exited.
func (s *BaseTransTotalListener) ExitTranstotal(ctx *TranstotalContext) {}
