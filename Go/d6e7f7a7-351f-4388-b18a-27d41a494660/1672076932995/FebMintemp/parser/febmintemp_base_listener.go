// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076932995/FebMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // FebMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFebMintempListener is a complete listener for a parse tree produced by FebMintempParser.
type BaseFebMintempListener struct{}

var _ FebMintempListener = &BaseFebMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFebMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFebMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFebMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFebMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFebMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFebMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFebmintemp is called when production febmintemp is entered.
func (s *BaseFebMintempListener) EnterFebmintemp(ctx *FebmintempContext) {}

// ExitFebmintemp is called when production febmintemp is exited.
func (s *BaseFebMintempListener) ExitFebmintemp(ctx *FebmintempContext) {}
