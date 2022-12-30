// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376765383/EngineVolume.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // EngineVolume

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEngineVolumeListener is a complete listener for a parse tree produced by EngineVolumeParser.
type BaseEngineVolumeListener struct{}

var _ EngineVolumeListener = &BaseEngineVolumeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEngineVolumeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEngineVolumeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEngineVolumeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEngineVolumeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEngineVolumeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEngineVolumeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEnginevolume is called when production enginevolume is entered.
func (s *BaseEngineVolumeListener) EnterEnginevolume(ctx *EnginevolumeContext) {}

// ExitEnginevolume is called when production enginevolume is exited.
func (s *BaseEngineVolumeListener) ExitEnginevolume(ctx *EnginevolumeContext) {}
