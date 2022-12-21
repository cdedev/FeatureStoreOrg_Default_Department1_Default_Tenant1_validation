// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671598824557/Coal.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Coal

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCoalListener is a complete listener for a parse tree produced by CoalParser.
type BaseCoalListener struct{}

var _ CoalListener = &BaseCoalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCoalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCoalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCoalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCoalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCoalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCoalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCoal is called when production coal is entered.
func (s *BaseCoalListener) EnterCoal(ctx *CoalContext) {}

// ExitCoal is called when production coal is exited.
func (s *BaseCoalListener) ExitCoal(ctx *CoalContext) {}
