// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669659170926/SquareFeet.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SquareFeet

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSquareFeetListener is a complete listener for a parse tree produced by SquareFeetParser.
type BaseSquareFeetListener struct{}

var _ SquareFeetListener = &BaseSquareFeetListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSquareFeetListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSquareFeetListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSquareFeetListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSquareFeetListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSquareFeetListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSquareFeetListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSquarefeet is called when production squarefeet is entered.
func (s *BaseSquareFeetListener) EnterSquarefeet(ctx *SquarefeetContext) {}

// ExitSquarefeet is called when production squarefeet is exited.
func (s *BaseSquareFeetListener) ExitSquarefeet(ctx *SquarefeetContext) {}
