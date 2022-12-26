// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077266801/JunMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JunMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJunMintempListener is a complete listener for a parse tree produced by JunMintempParser.
type BaseJunMintempListener struct{}

var _ JunMintempListener = &BaseJunMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJunMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJunMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJunMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJunMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJunMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJunMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJunmintemp is called when production junmintemp is entered.
func (s *BaseJunMintempListener) EnterJunmintemp(ctx *JunmintempContext) {}

// ExitJunmintemp is called when production junmintemp is exited.
func (s *BaseJunMintempListener) ExitJunmintemp(ctx *JunmintempContext) {}
