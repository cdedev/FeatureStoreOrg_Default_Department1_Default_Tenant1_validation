// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675306627938/Fat.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Fat

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFatListener is a complete listener for a parse tree produced by FatParser.
type BaseFatListener struct{}

var _ FatListener = &BaseFatListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFatListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFatListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFatListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFatListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFatListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFatListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFat is called when production fat is entered.
func (s *BaseFatListener) EnterFat(ctx *FatContext) {}

// ExitFat is called when production fat is exited.
func (s *BaseFatListener) ExitFat(ctx *FatContext) {}
