// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672120114803/Elevation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Elevation

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseElevationListener is a complete listener for a parse tree produced by ElevationParser.
type BaseElevationListener struct{}

var _ ElevationListener = &BaseElevationListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseElevationListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseElevationListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseElevationListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseElevationListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseElevationListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseElevationListener) ExitExpression(ctx *ExpressionContext) {}

// EnterElevation is called when production elevation is entered.
func (s *BaseElevationListener) EnterElevation(ctx *ElevationContext) {}

// ExitElevation is called when production elevation is exited.
func (s *BaseElevationListener) ExitElevation(ctx *ElevationContext) {}
