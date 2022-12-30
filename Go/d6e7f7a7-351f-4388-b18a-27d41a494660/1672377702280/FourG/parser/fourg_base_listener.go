// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672377702280/FourG.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FourG

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFourGListener is a complete listener for a parse tree produced by FourGParser.
type BaseFourGListener struct{}

var _ FourGListener = &BaseFourGListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFourGListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFourGListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFourGListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFourGListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFourGListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFourGListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFourg is called when production fourg is entered.
func (s *BaseFourGListener) EnterFourg(ctx *FourgContext) {}

// ExitFourg is called when production fourg is exited.
func (s *BaseFourGListener) ExitFourg(ctx *FourgContext) {}
