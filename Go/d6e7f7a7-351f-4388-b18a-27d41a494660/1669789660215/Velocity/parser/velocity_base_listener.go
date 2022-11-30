// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669789660215/Velocity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Velocity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVelocityListener is a complete listener for a parse tree produced by VelocityParser.
type BaseVelocityListener struct{}

var _ VelocityListener = &BaseVelocityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVelocityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVelocityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVelocityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVelocityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVelocityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVelocityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVelocity is called when production velocity is entered.
func (s *BaseVelocityListener) EnterVelocity(ctx *VelocityContext) {}

// ExitVelocity is called when production velocity is exited.
func (s *BaseVelocityListener) ExitVelocity(ctx *VelocityContext) {}
