// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077184093/MayMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MayMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMayMintempListener is a complete listener for a parse tree produced by MayMintempParser.
type BaseMayMintempListener struct{}

var _ MayMintempListener = &BaseMayMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMayMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMayMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMayMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMayMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMayMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMayMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMaymintemp is called when production maymintemp is entered.
func (s *BaseMayMintempListener) EnterMaymintemp(ctx *MaymintempContext) {}

// ExitMaymintemp is called when production maymintemp is exited.
func (s *BaseMayMintempListener) ExitMaymintemp(ctx *MaymintempContext) {}
