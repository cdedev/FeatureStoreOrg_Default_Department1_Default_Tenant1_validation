// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078015413/NovMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NovMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNovMintempListener is a complete listener for a parse tree produced by NovMintempParser.
type BaseNovMintempListener struct{}

var _ NovMintempListener = &BaseNovMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNovMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNovMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNovMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNovMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNovMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNovMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNovmintemp is called when production novmintemp is entered.
func (s *BaseNovMintempListener) EnterNovmintemp(ctx *NovmintempContext) {}

// ExitNovmintemp is called when production novmintemp is exited.
func (s *BaseNovMintempListener) ExitNovmintemp(ctx *NovmintempContext) {}
