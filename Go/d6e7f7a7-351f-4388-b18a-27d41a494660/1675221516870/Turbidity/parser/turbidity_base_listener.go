// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675221516870/Turbidity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Turbidity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTurbidityListener is a complete listener for a parse tree produced by TurbidityParser.
type BaseTurbidityListener struct{}

var _ TurbidityListener = &BaseTurbidityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTurbidityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTurbidityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTurbidityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTurbidityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTurbidityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTurbidityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTurbidity is called when production turbidity is entered.
func (s *BaseTurbidityListener) EnterTurbidity(ctx *TurbidityContext) {}

// ExitTurbidity is called when production turbidity is exited.
func (s *BaseTurbidityListener) ExitTurbidity(ctx *TurbidityContext) {}
