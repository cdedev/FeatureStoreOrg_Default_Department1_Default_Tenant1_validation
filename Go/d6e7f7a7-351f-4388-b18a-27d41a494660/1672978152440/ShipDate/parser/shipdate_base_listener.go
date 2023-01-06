// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672978152440/ShipDate.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ShipDate

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseShipDateListener is a complete listener for a parse tree produced by ShipDateParser.
type BaseShipDateListener struct{}

var _ ShipDateListener = &BaseShipDateListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseShipDateListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseShipDateListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseShipDateListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseShipDateListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseShipDateListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseShipDateListener) ExitExpression(ctx *ExpressionContext) {}

// EnterShipdate is called when production shipdate is entered.
func (s *BaseShipDateListener) EnterShipdate(ctx *ShipdateContext) {}

// ExitShipdate is called when production shipdate is exited.
func (s *BaseShipDateListener) ExitShipdate(ctx *ShipdateContext) {}
