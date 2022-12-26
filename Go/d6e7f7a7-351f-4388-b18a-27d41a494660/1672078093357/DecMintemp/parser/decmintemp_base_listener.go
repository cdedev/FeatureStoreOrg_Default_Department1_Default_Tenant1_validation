// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078093357/DecMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DecMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDecMintempListener is a complete listener for a parse tree produced by DecMintempParser.
type BaseDecMintempListener struct{}

var _ DecMintempListener = &BaseDecMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDecMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDecMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDecMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDecMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDecMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDecMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDecmintemp is called when production decmintemp is entered.
func (s *BaseDecMintempListener) EnterDecmintemp(ctx *DecmintempContext) {}

// ExitDecmintemp is called when production decmintemp is exited.
func (s *BaseDecMintempListener) ExitDecmintemp(ctx *DecmintempContext) {}
