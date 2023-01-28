// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674884745703/Town.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Town

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTownListener is a complete listener for a parse tree produced by TownParser.
type BaseTownListener struct{}

var _ TownListener = &BaseTownListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTownListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTownListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTownListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTownListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTownListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTownListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTown is called when production town is entered.
func (s *BaseTownListener) EnterTown(ctx *TownContext) {}

// ExitTown is called when production town is exited.
func (s *BaseTownListener) ExitTown(ctx *TownContext) {}
