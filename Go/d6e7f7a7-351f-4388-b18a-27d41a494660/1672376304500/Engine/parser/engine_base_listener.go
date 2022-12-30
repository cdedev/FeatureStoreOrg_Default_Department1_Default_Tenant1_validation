// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376304500/Engine.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Engine

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEngineListener is a complete listener for a parse tree produced by EngineParser.
type BaseEngineListener struct{}

var _ EngineListener = &BaseEngineListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEngineListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEngineListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEngineListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEngineListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEngineListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEngineListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEngine is called when production engine is entered.
func (s *BaseEngineListener) EnterEngine(ctx *EngineContext) {}

// ExitEngine is called when production engine is exited.
func (s *BaseEngineListener) ExitEngine(ctx *EngineContext) {}
