// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674447401032/Center.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Center

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCenterListener is a complete listener for a parse tree produced by CenterParser.
type BaseCenterListener struct{}

var _ CenterListener = &BaseCenterListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCenterListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCenterListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCenterListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCenterListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCenterListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCenterListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCenter is called when production center is entered.
func (s *BaseCenterListener) EnterCenter(ctx *CenterContext) {}

// ExitCenter is called when production center is exited.
func (s *BaseCenterListener) ExitCenter(ctx *CenterContext) {}
