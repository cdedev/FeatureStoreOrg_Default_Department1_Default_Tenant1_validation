// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673238889397/Sensitive.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Sensitive

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSensitiveListener is a complete listener for a parse tree produced by SensitiveParser.
type BaseSensitiveListener struct{}

var _ SensitiveListener = &BaseSensitiveListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSensitiveListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSensitiveListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSensitiveListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSensitiveListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSensitiveListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSensitiveListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSensitive is called when production sensitive is entered.
func (s *BaseSensitiveListener) EnterSensitive(ctx *SensitiveContext) {}

// ExitSensitive is called when production sensitive is exited.
func (s *BaseSensitiveListener) ExitSensitive(ctx *SensitiveContext) {}
