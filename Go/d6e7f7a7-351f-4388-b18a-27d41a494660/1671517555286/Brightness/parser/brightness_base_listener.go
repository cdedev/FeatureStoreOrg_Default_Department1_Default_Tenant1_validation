// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671517555286/Brightness.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Brightness

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBrightnessListener is a complete listener for a parse tree produced by BrightnessParser.
type BaseBrightnessListener struct{}

var _ BrightnessListener = &BaseBrightnessListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBrightnessListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBrightnessListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBrightnessListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBrightnessListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBrightnessListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBrightnessListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBrightness is called when production brightness is entered.
func (s *BaseBrightnessListener) EnterBrightness(ctx *BrightnessContext) {}

// ExitBrightness is called when production brightness is exited.
func (s *BaseBrightnessListener) ExitBrightness(ctx *BrightnessContext) {}
