// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077146818/MayMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MayMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMayMaxtempListener is a complete listener for a parse tree produced by MayMaxtempParser.
type BaseMayMaxtempListener struct{}

var _ MayMaxtempListener = &BaseMayMaxtempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMayMaxtempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMayMaxtempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMayMaxtempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMayMaxtempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMayMaxtempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMayMaxtempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMaymaxtemp is called when production maymaxtemp is entered.
func (s *BaseMayMaxtempListener) EnterMaymaxtemp(ctx *MaymaxtempContext) {}

// ExitMaymaxtemp is called when production maymaxtemp is exited.
func (s *BaseMayMaxtempListener) ExitMaymaxtemp(ctx *MaymaxtempContext) {}
