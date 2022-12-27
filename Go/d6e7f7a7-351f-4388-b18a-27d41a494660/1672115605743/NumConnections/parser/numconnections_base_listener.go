// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672115605743/NumConnections.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // NumConnections

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNumConnectionsListener is a complete listener for a parse tree produced by NumConnectionsParser.
type BaseNumConnectionsListener struct{}

var _ NumConnectionsListener = &BaseNumConnectionsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNumConnectionsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNumConnectionsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNumConnectionsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNumConnectionsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNumConnectionsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNumConnectionsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNumconnections is called when production numconnections is entered.
func (s *BaseNumConnectionsListener) EnterNumconnections(ctx *NumconnectionsContext) {}

// ExitNumconnections is called when production numconnections is exited.
func (s *BaseNumConnectionsListener) ExitNumconnections(ctx *NumconnectionsContext) {}
