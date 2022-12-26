// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672078049003/DecMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DecMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDecMaxtempListener is a complete listener for a parse tree produced by DecMaxtempParser.
type BaseDecMaxtempListener struct{}

var _ DecMaxtempListener = &BaseDecMaxtempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDecMaxtempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDecMaxtempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDecMaxtempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDecMaxtempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDecMaxtempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDecMaxtempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDecmaxtemp is called when production decmaxtemp is entered.
func (s *BaseDecMaxtempListener) EnterDecmaxtemp(ctx *DecmaxtempContext) {}

// ExitDecmaxtemp is called when production decmaxtemp is exited.
func (s *BaseDecMaxtempListener) ExitDecmaxtemp(ctx *DecmaxtempContext) {}
