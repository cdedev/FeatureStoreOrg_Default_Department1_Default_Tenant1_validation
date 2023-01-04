// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672810323495/Circumference.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Circumference

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCircumferenceListener is a complete listener for a parse tree produced by CircumferenceParser.
type BaseCircumferenceListener struct{}

var _ CircumferenceListener = &BaseCircumferenceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCircumferenceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCircumferenceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCircumferenceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCircumferenceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCircumferenceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCircumferenceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCircumference is called when production circumference is entered.
func (s *BaseCircumferenceListener) EnterCircumference(ctx *CircumferenceContext) {}

// ExitCircumference is called when production circumference is exited.
func (s *BaseCircumferenceListener) ExitCircumference(ctx *CircumferenceContext) {}
