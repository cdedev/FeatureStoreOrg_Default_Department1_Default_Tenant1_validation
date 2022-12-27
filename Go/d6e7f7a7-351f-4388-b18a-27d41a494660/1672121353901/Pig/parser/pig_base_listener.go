// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672121353901/Pig.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Pig

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePigListener is a complete listener for a parse tree produced by PigParser.
type BasePigListener struct{}

var _ PigListener = &BasePigListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePigListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePigListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePigListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePigListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePigListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePigListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPig is called when production pig is entered.
func (s *BasePigListener) EnterPig(ctx *PigContext) {}

// ExitPig is called when production pig is exited.
func (s *BasePigListener) ExitPig(ctx *PigContext) {}
