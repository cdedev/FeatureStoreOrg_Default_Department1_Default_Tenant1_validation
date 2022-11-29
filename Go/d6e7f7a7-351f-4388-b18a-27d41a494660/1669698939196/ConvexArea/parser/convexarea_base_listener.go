// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669698939196/ConvexArea.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ConvexArea

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseConvexAreaListener is a complete listener for a parse tree produced by ConvexAreaParser.
type BaseConvexAreaListener struct{}

var _ ConvexAreaListener = &BaseConvexAreaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseConvexAreaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseConvexAreaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseConvexAreaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseConvexAreaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseConvexAreaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseConvexAreaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConvexarea is called when production convexarea is entered.
func (s *BaseConvexAreaListener) EnterConvexarea(ctx *ConvexareaContext) {}

// ExitConvexarea is called when production convexarea is exited.
func (s *BaseConvexAreaListener) ExitConvexarea(ctx *ConvexareaContext) {}
