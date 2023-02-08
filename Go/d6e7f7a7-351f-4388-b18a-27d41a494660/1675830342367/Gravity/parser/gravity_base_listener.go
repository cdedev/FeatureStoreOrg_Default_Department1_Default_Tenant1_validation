// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675830342367/Gravity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Gravity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGravityListener is a complete listener for a parse tree produced by GravityParser.
type BaseGravityListener struct{}

var _ GravityListener = &BaseGravityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGravityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGravityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGravityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGravityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGravityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGravityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGravity is called when production gravity is entered.
func (s *BaseGravityListener) EnterGravity(ctx *GravityContext) {}

// ExitGravity is called when production gravity is exited.
func (s *BaseGravityListener) ExitGravity(ctx *GravityContext) {}
