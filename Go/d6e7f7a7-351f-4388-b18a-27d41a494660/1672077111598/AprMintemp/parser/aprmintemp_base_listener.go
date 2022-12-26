// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077111598/AprMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AprMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAprMintempListener is a complete listener for a parse tree produced by AprMintempParser.
type BaseAprMintempListener struct{}

var _ AprMintempListener = &BaseAprMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAprMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAprMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAprMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAprMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAprMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAprMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAprmintemp is called when production aprmintemp is entered.
func (s *BaseAprMintempListener) EnterAprmintemp(ctx *AprmintempContext) {}

// ExitAprmintemp is called when production aprmintemp is exited.
func (s *BaseAprMintempListener) ExitAprmintemp(ctx *AprmintempContext) {}
