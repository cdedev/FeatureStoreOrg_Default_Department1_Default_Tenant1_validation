// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671694797816/Network.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Network

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNetworkListener is a complete listener for a parse tree produced by NetworkParser.
type BaseNetworkListener struct{}

var _ NetworkListener = &BaseNetworkListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNetworkListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNetworkListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNetworkListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNetworkListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNetworkListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNetworkListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNetwork is called when production network is entered.
func (s *BaseNetworkListener) EnterNetwork(ctx *NetworkContext) {}

// ExitNetwork is called when production network is exited.
func (s *BaseNetworkListener) ExitNetwork(ctx *NetworkContext) {}
