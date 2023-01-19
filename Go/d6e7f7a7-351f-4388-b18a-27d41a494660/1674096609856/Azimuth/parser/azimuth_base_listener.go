// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674096609856/Azimuth.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Azimuth

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAzimuthListener is a complete listener for a parse tree produced by AzimuthParser.
type BaseAzimuthListener struct{}

var _ AzimuthListener = &BaseAzimuthListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAzimuthListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAzimuthListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAzimuthListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAzimuthListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseAzimuthListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseAzimuthListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAzimuth is called when production azimuth is entered.
func (s *BaseAzimuthListener) EnterAzimuth(ctx *AzimuthContext) {}

// ExitAzimuth is called when production azimuth is exited.
func (s *BaseAzimuthListener) ExitAzimuth(ctx *AzimuthContext) {}
