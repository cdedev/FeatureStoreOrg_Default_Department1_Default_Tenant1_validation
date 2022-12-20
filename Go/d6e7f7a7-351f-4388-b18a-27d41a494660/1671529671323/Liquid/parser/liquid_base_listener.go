// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671529671323/Liquid.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Liquid

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLiquidListener is a complete listener for a parse tree produced by LiquidParser.
type BaseLiquidListener struct{}

var _ LiquidListener = &BaseLiquidListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLiquidListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLiquidListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLiquidListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLiquidListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLiquidListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLiquidListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLiquid is called when production liquid is entered.
func (s *BaseLiquidListener) EnterLiquid(ctx *LiquidContext) {}

// ExitLiquid is called when production liquid is exited.
func (s *BaseLiquidListener) ExitLiquid(ctx *LiquidContext) {}
