// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672284909159/Nudity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Nudity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNudityListener is a complete listener for a parse tree produced by NudityParser.
type BaseNudityListener struct{}

var _ NudityListener = &BaseNudityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNudityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNudityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNudityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNudityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseNudityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseNudityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterNudity is called when production nudity is entered.
func (s *BaseNudityListener) EnterNudity(ctx *NudityContext) {}

// ExitNudity is called when production nudity is exited.
func (s *BaseNudityListener) ExitNudity(ctx *NudityContext) {}
