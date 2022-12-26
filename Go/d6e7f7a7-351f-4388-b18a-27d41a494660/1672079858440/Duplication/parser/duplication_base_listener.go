// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672079858440/Duplication.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Duplication

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDuplicationListener is a complete listener for a parse tree produced by DuplicationParser.
type BaseDuplicationListener struct{}

var _ DuplicationListener = &BaseDuplicationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDuplicationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDuplicationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDuplicationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDuplicationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDuplicationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDuplicationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDuplication is called when production duplication is entered.
func (s *BaseDuplicationListener) EnterDuplication(ctx *DuplicationContext) {}

// ExitDuplication is called when production duplication is exited.
func (s *BaseDuplicationListener) ExitDuplication(ctx *DuplicationContext) {}
