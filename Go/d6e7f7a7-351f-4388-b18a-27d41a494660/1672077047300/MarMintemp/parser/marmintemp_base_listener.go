// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077047300/MarMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MarMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMarMintempListener is a complete listener for a parse tree produced by MarMintempParser.
type BaseMarMintempListener struct{}

var _ MarMintempListener = &BaseMarMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMarMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMarMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMarmintemp is called when production marmintemp is entered.
func (s *BaseMarMintempListener) EnterMarmintemp(ctx *MarmintempContext) {}

// ExitMarmintemp is called when production marmintemp is exited.
func (s *BaseMarMintempListener) ExitMarmintemp(ctx *MarmintempContext) {}
