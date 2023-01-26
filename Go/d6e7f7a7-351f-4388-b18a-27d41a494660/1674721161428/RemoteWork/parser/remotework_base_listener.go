// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674721161428/RemoteWork.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // RemoteWork

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseRemoteWorkListener is a complete listener for a parse tree produced by RemoteWorkParser.
type BaseRemoteWorkListener struct{}

var _ RemoteWorkListener = &BaseRemoteWorkListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRemoteWorkListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRemoteWorkListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRemoteWorkListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRemoteWorkListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseRemoteWorkListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseRemoteWorkListener) ExitExpression(ctx *ExpressionContext) {}

// EnterRemotework is called when production remotework is entered.
func (s *BaseRemoteWorkListener) EnterRemotework(ctx *RemoteworkContext) {}

// ExitRemotework is called when production remotework is exited.
func (s *BaseRemoteWorkListener) ExitRemotework(ctx *RemoteworkContext) {}
