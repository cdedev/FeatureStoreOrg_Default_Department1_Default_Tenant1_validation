// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669778108485/DestinationPort.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // DestinationPort

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDestinationPortListener is a complete listener for a parse tree produced by DestinationPortParser.
type BaseDestinationPortListener struct{}

var _ DestinationPortListener = &BaseDestinationPortListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDestinationPortListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDestinationPortListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDestinationPortListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDestinationPortListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDestinationPortListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDestinationPortListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDestinationport is called when production destinationport is entered.
func (s *BaseDestinationPortListener) EnterDestinationport(ctx *DestinationportContext) {}

// ExitDestinationport is called when production destinationport is exited.
func (s *BaseDestinationPortListener) ExitDestinationport(ctx *DestinationportContext) {}
