// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675931175351/Nauseavomit.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nauseavomit

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNauseavomitListener is a complete listener for a parse tree produced by NauseavomitParser.
type BaseNauseavomitListener struct{}

var _ NauseavomitListener = &BaseNauseavomitListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNauseavomitListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNauseavomitListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNauseavomitListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNauseavomitListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNauseavomitListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNauseavomitListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNauseavomit is called when production nauseavomit is entered.
func (s *BaseNauseavomitListener) EnterNauseavomit(ctx *NauseavomitContext) {}

// ExitNauseavomit is called when production nauseavomit is exited.
func (s *BaseNauseavomitListener) ExitNauseavomit(ctx *NauseavomitContext) {}
