// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1672376810760/GearBoxType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // GearBoxType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGearBoxTypeListener is a complete listener for a parse tree produced by GearBoxTypeParser.
type BaseGearBoxTypeListener struct{}

var _ GearBoxTypeListener = &BaseGearBoxTypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGearBoxTypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGearBoxTypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGearBoxTypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGearBoxTypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGearBoxTypeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGearBoxTypeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterGearboxtype is called when production gearboxtype is entered.
func (s *BaseGearBoxTypeListener) EnterGearboxtype(ctx *GearboxtypeContext) {}

// ExitGearboxtype is called when production gearboxtype is exited.
func (s *BaseGearBoxTypeListener) ExitGearboxtype(ctx *GearboxtypeContext) {}
