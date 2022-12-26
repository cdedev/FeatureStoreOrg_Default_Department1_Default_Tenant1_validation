// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077506509/AugMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // AugMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAugMaxtempListener is a complete listener for a parse tree produced by AugMaxtempParser.
type BaseAugMaxtempListener struct{}

var _ AugMaxtempListener = &BaseAugMaxtempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAugMaxtempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAugMaxtempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAugMaxtempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAugMaxtempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAugMaxtempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAugMaxtempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAugmaxtemp is called when production augmaxtemp is entered.
func (s *BaseAugMaxtempListener) EnterAugmaxtemp(ctx *AugmaxtempContext) {}

// ExitAugmaxtemp is called when production augmaxtemp is exited.
func (s *BaseAugMaxtempListener) ExitAugmaxtemp(ctx *AugmaxtempContext) {}
