// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672077225099/JunMaxtemp.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // JunMaxtemp

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseJunMaxtempListener is a complete listener for a parse tree produced by JunMaxtempParser.
type BaseJunMaxtempListener struct{}

var _ JunMaxtempListener = &BaseJunMaxtempListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJunMaxtempListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJunMaxtempListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJunMaxtempListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJunMaxtempListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseJunMaxtempListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseJunMaxtempListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJunmaxtemp is called when production junmaxtemp is entered.
func (s *BaseJunMaxtempListener) EnterJunmaxtemp(ctx *JunmaxtempContext) {}

// ExitJunmaxtemp is called when production junmaxtemp is exited.
func (s *BaseJunMaxtempListener) ExitJunmaxtemp(ctx *JunmaxtempContext) {}
