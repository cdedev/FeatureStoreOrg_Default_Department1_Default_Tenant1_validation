// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669697837214/Chlorides.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Chlorides

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChloridesListener is a complete listener for a parse tree produced by ChloridesParser.
type BaseChloridesListener struct{}

var _ ChloridesListener = &BaseChloridesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChloridesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChloridesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChloridesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChloridesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChloridesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChloridesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterChlorides is called when production chlorides is entered.
func (s *BaseChloridesListener) EnterChlorides(ctx *ChloridesContext) {}

// ExitChlorides is called when production chlorides is exited.
func (s *BaseChloridesListener) ExitChlorides(ctx *ChloridesContext) {}
