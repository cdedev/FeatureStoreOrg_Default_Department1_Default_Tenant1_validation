// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077895049/OctMintemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // OctMintemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseOctMintempListener is a complete listener for a parse tree produced by OctMintempParser.
type BaseOctMintempListener struct{}

var _ OctMintempListener = &BaseOctMintempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOctMintempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOctMintempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOctMintempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOctMintempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseOctMintempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseOctMintempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterOctmintemp is called when production octmintemp is entered.
func (s *BaseOctMintempListener) EnterOctmintemp(ctx *OctmintempContext) {}

// ExitOctmintemp is called when production octmintemp is exited.
func (s *BaseOctMintempListener) ExitOctmintemp(ctx *OctmintempContext) {}
