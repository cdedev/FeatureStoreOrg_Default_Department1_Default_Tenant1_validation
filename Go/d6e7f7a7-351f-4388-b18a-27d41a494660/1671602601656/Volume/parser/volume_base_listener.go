// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671602601656/Volume.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Volume

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVolumeListener is a complete listener for a parse tree produced by VolumeParser.
type BaseVolumeListener struct{}

var _ VolumeListener = &BaseVolumeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVolumeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVolumeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVolumeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVolumeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVolumeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVolumeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVolume is called when production volume is entered.
func (s *BaseVolumeListener) EnterVolume(ctx *VolumeContext) {}

// ExitVolume is called when production volume is exited.
func (s *BaseVolumeListener) ExitVolume(ctx *VolumeContext) {}
