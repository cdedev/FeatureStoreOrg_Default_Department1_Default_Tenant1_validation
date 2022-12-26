// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077628694/SepMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SepMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSepMaxtempListener is a complete listener for a parse tree produced by SepMaxtempParser.
type BaseSepMaxtempListener struct{}

var _ SepMaxtempListener = &BaseSepMaxtempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSepMaxtempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSepMaxtempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSepMaxtempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSepMaxtempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSepMaxtempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSepMaxtempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterSepmaxtemp is called when production sepmaxtemp is entered.
func (s *BaseSepMaxtempListener) EnterSepmaxtemp(ctx *SepmaxtempContext) {}

// ExitSepmaxtemp is called when production sepmaxtemp is exited.
func (s *BaseSepMaxtempListener) ExitSepmaxtemp(ctx *SepmaxtempContext) {}
