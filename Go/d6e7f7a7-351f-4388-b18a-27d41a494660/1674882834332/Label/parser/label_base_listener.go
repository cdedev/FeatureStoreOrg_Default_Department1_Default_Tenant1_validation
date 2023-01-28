// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674882834332/Label.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Label

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLabelListener is a complete listener for a parse tree produced by LabelParser.
type BaseLabelListener struct{}

var _ LabelListener = &BaseLabelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLabelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLabelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLabelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLabelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLabelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLabelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseLabelListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseLabelListener) ExitLabel(ctx *LabelContext) {}
