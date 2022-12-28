// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672204772426/Lightbill.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Lightbill

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLightbillListener is a complete listener for a parse tree produced by LightbillParser.
type BaseLightbillListener struct{}

var _ LightbillListener = &BaseLightbillListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLightbillListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLightbillListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLightbillListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLightbillListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLightbillListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLightbillListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLightbill is called when production lightbill is entered.
func (s *BaseLightbillListener) EnterLightbill(ctx *LightbillContext) {}

// ExitLightbill is called when production lightbill is exited.
func (s *BaseLightbillListener) ExitLightbill(ctx *LightbillContext) {}
