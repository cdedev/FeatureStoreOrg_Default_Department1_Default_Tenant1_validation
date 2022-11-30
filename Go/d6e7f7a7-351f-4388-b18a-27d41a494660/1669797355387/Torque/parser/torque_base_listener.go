// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669797355387/Torque.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Torque

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTorqueListener is a complete listener for a parse tree produced by TorqueParser.
type BaseTorqueListener struct{}

var _ TorqueListener = &BaseTorqueListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTorqueListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTorqueListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTorqueListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTorqueListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTorqueListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTorqueListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTorque is called when production torque is entered.
func (s *BaseTorqueListener) EnterTorque(ctx *TorqueContext) {}

// ExitTorque is called when production torque is exited.
func (s *BaseTorqueListener) ExitTorque(ctx *TorqueContext) {}
