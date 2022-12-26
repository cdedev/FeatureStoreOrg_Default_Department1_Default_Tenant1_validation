// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077012015/MarMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MarMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMarMaxtempListener is a complete listener for a parse tree produced by MarMaxtempParser.
type BaseMarMaxtempListener struct{}

var _ MarMaxtempListener = &BaseMarMaxtempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMarMaxtempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMarMaxtempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMarMaxtempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMarMaxtempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseMarMaxtempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseMarMaxtempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterMarmaxtemp is called when production marmaxtemp is entered.
func (s *BaseMarMaxtempListener) EnterMarmaxtemp(ctx *MarmaxtempContext) {}

// ExitMarmaxtemp is called when production marmaxtemp is exited.
func (s *BaseMarMaxtempListener) ExitMarmaxtemp(ctx *MarmaxtempContext) {}
