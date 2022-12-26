// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077553967/AugMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AugMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAugMintempListener is a complete listener for a parse tree produced by AugMintempParser.
type BaseAugMintempListener struct{}

var _ AugMintempListener = &BaseAugMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAugMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAugMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAugMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAugMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAugMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAugMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAugmintemp is called when production augmintemp is entered.
func (s *BaseAugMintempListener) EnterAugmintemp(ctx *AugmintempContext) {}

// ExitAugmintemp is called when production augmintemp is exited.
func (s *BaseAugMintempListener) ExitAugmintemp(ctx *AugmintempContext) {}
