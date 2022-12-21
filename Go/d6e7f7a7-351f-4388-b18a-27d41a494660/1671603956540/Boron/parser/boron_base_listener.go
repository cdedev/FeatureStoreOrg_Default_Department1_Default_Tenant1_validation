// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671603956540/Boron.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Boron

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBoronListener is a complete listener for a parse tree produced by BoronParser.
type BaseBoronListener struct{}

var _ BoronListener = &BaseBoronListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBoronListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBoronListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBoronListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBoronListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBoronListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBoronListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBoron is called when production boron is entered.
func (s *BaseBoronListener) EnterBoron(ctx *BoronContext) {}

// ExitBoron is called when production boron is exited.
func (s *BaseBoronListener) ExitBoron(ctx *BoronContext) {}
