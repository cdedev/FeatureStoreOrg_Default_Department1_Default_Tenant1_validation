// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515379219/VehicleType.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // VehicleType

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseVehicleTypeListener is a complete listener for a parse tree produced by VehicleTypeParser.
type BaseVehicleTypeListener struct{}

var _ VehicleTypeListener = &BaseVehicleTypeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVehicleTypeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVehicleTypeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVehicleTypeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVehicleTypeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseVehicleTypeListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseVehicleTypeListener) ExitExpression(ctx *ExpressionContext) {}

// EnterVehicletype is called when production vehicletype is entered.
func (s *BaseVehicleTypeListener) EnterVehicletype(ctx *VehicletypeContext) {}

// ExitVehicletype is called when production vehicletype is exited.
func (s *BaseVehicleTypeListener) ExitVehicletype(ctx *VehicletypeContext) {}
