// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676606932570/HairLoss.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // HairLoss

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseHairLossListener is a complete listener for a parse tree produced by HairLossParser.
type BaseHairLossListener struct{}

var _ HairLossListener = &BaseHairLossListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHairLossListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHairLossListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHairLossListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHairLossListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseHairLossListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseHairLossListener) ExitExpression(ctx *ExpressionContext) {}

// EnterHairloss is called when production hairloss is entered.
func (s *BaseHairLossListener) EnterHairloss(ctx *HairlossContext) {}

// ExitHairloss is called when production hairloss is exited.
func (s *BaseHairLossListener) ExitHairloss(ctx *HairlossContext) {}
