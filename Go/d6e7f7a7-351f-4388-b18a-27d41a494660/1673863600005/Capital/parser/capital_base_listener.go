// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673863600005/Capital.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Capital

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCapitalListener is a complete listener for a parse tree produced by CapitalParser.
type BaseCapitalListener struct{}

var _ CapitalListener = &BaseCapitalListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCapitalListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCapitalListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCapitalListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCapitalListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseCapitalListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCapitalListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCapital is called when production capital is entered.
func (s *BaseCapitalListener) EnterCapital(ctx *CapitalContext) {}

// ExitCapital is called when production capital is exited.
func (s *BaseCapitalListener) ExitCapital(ctx *CapitalContext) {}
