// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377355661/TouchScreen.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // TouchScreen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTouchScreenListener is a complete listener for a parse tree produced by TouchScreenParser.
type BaseTouchScreenListener struct{}

var _ TouchScreenListener = &BaseTouchScreenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTouchScreenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTouchScreenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTouchScreenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTouchScreenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTouchScreenListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTouchScreenListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTouchscreen is called when production touchscreen is entered.
func (s *BaseTouchScreenListener) EnterTouchscreen(ctx *TouchscreenContext) {}

// ExitTouchscreen is called when production touchscreen is exited.
func (s *BaseTouchScreenListener) ExitTouchscreen(ctx *TouchscreenContext) {}
