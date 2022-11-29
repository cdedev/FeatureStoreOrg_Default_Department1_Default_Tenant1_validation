// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697774224/Alcohol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Alcohol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAlcoholListener is a complete listener for a parse tree produced by AlcoholParser.
type BaseAlcoholListener struct{}

var _ AlcoholListener = &BaseAlcoholListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAlcoholListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAlcoholListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAlcoholListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAlcoholListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAlcoholListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAlcoholListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAlcohol is called when production alcohol is entered.
func (s *BaseAlcoholListener) EnterAlcohol(ctx *AlcoholContext) {}

// ExitAlcohol is called when production alcohol is exited.
func (s *BaseAlcoholListener) ExitAlcohol(ctx *AlcoholContext) {}
