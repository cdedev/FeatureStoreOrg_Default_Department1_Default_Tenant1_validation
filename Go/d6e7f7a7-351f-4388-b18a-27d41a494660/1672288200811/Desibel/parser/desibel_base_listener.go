// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672288200811/Desibel.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Desibel

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDesibelListener is a complete listener for a parse tree produced by DesibelParser.
type BaseDesibelListener struct{}

var _ DesibelListener = &BaseDesibelListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDesibelListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDesibelListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDesibelListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDesibelListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDesibelListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDesibelListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDesibel is called when production desibel is entered.
func (s *BaseDesibelListener) EnterDesibel(ctx *DesibelContext) {}

// ExitDesibel is called when production desibel is exited.
func (s *BaseDesibelListener) ExitDesibel(ctx *DesibelContext) {}
