// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669699123294/Eccentricity.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Eccentricity

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseEccentricityListener is a complete listener for a parse tree produced by EccentricityParser.
type BaseEccentricityListener struct{}

var _ EccentricityListener = &BaseEccentricityListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEccentricityListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEccentricityListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEccentricityListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEccentricityListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseEccentricityListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseEccentricityListener) ExitExpression(ctx *ExpressionContext) {}

// EnterEccentricity is called when production eccentricity is entered.
func (s *BaseEccentricityListener) EnterEccentricity(ctx *EccentricityContext) {}

// ExitEccentricity is called when production eccentricity is exited.
func (s *BaseEccentricityListener) ExitEccentricity(ctx *EccentricityContext) {}
