// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675833307329/Cholesterol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cholesterol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCholesterolListener is a complete listener for a parse tree produced by CholesterolParser.
type BaseCholesterolListener struct{}

var _ CholesterolListener = &BaseCholesterolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCholesterolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCholesterolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCholesterolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCholesterolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCholesterolListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCholesterolListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCholesterol is called when production cholesterol is entered.
func (s *BaseCholesterolListener) EnterCholesterol(ctx *CholesterolContext) {}

// ExitCholesterol is called when production cholesterol is exited.
func (s *BaseCholesterolListener) ExitCholesterol(ctx *CholesterolContext) {}
