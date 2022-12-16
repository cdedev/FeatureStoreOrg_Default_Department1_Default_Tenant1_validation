// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671202914585/Greenzones.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Greenzones

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGreenzonesListener is a complete listener for a parse tree produced by GreenzonesParser.
type BaseGreenzonesListener struct{}

var _ GreenzonesListener = &BaseGreenzonesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGreenzonesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGreenzonesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGreenzonesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGreenzonesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGreenzonesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGreenzonesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGreenzones is called when production greenzones is entered.
func (s *BaseGreenzonesListener) EnterGreenzones(ctx *GreenzonesContext) {}

// ExitGreenzones is called when production greenzones is exited.
func (s *BaseGreenzonesListener) ExitGreenzones(ctx *GreenzonesContext) {}
