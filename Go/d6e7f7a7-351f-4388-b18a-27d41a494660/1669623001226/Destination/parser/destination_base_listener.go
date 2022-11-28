// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669623001226/Destination.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Destination

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDestinationListener is a complete listener for a parse tree produced by DestinationParser.
type BaseDestinationListener struct{}

var _ DestinationListener = &BaseDestinationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDestinationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDestinationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDestinationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDestinationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDestinationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDestinationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDestination is called when production destination is entered.
func (s *BaseDestinationListener) EnterDestination(ctx *DestinationContext) {}

// ExitDestination is called when production destination is exited.
func (s *BaseDestinationListener) ExitDestination(ctx *DestinationContext) {}
