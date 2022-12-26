// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672076825716/JanMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JanMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJanMaxtempListener is a complete listener for a parse tree produced by JanMaxtempParser.
type BaseJanMaxtempListener struct{}

var _ JanMaxtempListener = &BaseJanMaxtempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJanMaxtempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJanMaxtempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJanMaxtempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJanMaxtempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJanMaxtempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJanMaxtempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJanmaxtemp is called when production janmaxtemp is entered.
func (s *BaseJanMaxtempListener) EnterJanmaxtemp(ctx *JanmaxtempContext) {}

// ExitJanmaxtemp is called when production janmaxtemp is exited.
func (s *BaseJanMaxtempListener) ExitJanmaxtemp(ctx *JanmaxtempContext) {}
