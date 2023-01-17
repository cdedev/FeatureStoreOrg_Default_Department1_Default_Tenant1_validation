// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1673925038138/Backbone.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Backbone

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBackboneListener is a complete listener for a parse tree produced by BackboneParser.
type BaseBackboneListener struct{}

var _ BackboneListener = &BaseBackboneListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBackboneListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBackboneListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBackboneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBackboneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBackboneListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBackboneListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBackbone is called when production backbone is entered.
func (s *BaseBackboneListener) EnterBackbone(ctx *BackboneContext) {}

// ExitBackbone is called when production backbone is exited.
func (s *BaseBackboneListener) ExitBackbone(ctx *BackboneContext) {}
