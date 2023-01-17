// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925550783/Fins.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fins

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFinsListener is a complete listener for a parse tree produced by FinsParser.
type BaseFinsListener struct{}

var _ FinsListener = &BaseFinsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFinsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFinsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFinsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFinsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFinsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFinsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFins is called when production fins is entered.
func (s *BaseFinsListener) EnterFins(ctx *FinsContext) {}

// ExitFins is called when production fins is exited.
func (s *BaseFinsListener) ExitFins(ctx *FinsContext) {}
