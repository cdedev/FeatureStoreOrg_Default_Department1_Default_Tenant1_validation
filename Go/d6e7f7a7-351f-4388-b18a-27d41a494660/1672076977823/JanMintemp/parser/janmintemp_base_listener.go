// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076977823/JanMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JanMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJanMintempListener is a complete listener for a parse tree produced by JanMintempParser.
type BaseJanMintempListener struct{}

var _ JanMintempListener = &BaseJanMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJanMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJanMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJanMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJanMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJanMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJanMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJanmintemp is called when production janmintemp is entered.
func (s *BaseJanMintempListener) EnterJanmintemp(ctx *JanmintempContext) {}

// ExitJanmintemp is called when production janmintemp is exited.
func (s *BaseJanMintempListener) ExitJanmintemp(ctx *JanmintempContext) {}
