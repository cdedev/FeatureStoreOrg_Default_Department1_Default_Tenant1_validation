// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672380066598/Fever.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fever

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFeverListener is a complete listener for a parse tree produced by FeverParser.
type BaseFeverListener struct{}

var _ FeverListener = &BaseFeverListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFeverListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFeverListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFeverListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFeverListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFeverListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFeverListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFever is called when production fever is entered.
func (s *BaseFeverListener) EnterFever(ctx *FeverContext) {}

// ExitFever is called when production fever is exited.
func (s *BaseFeverListener) ExitFever(ctx *FeverContext) {}
