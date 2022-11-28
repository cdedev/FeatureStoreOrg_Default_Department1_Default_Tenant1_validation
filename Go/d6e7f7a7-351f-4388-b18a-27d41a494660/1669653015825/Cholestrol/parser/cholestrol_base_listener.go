// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669653015825/Cholestrol.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Cholestrol

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCholestrolListener is a complete listener for a parse tree produced by CholestrolParser.
type BaseCholestrolListener struct{}

var _ CholestrolListener = &BaseCholestrolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCholestrolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCholestrolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCholestrolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCholestrolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCholestrolListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCholestrolListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCholestrol is called when production cholestrol is entered.
func (s *BaseCholestrolListener) EnterCholestrol(ctx *CholestrolContext) {}

// ExitCholestrol is called when production cholestrol is exited.
func (s *BaseCholestrolListener) ExitCholestrol(ctx *CholestrolContext) {}
