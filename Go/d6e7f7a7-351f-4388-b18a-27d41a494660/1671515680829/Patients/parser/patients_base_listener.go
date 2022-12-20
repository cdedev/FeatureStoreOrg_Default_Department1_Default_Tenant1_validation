// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1671515680829/Patients.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Patients

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePatientsListener is a complete listener for a parse tree produced by PatientsParser.
type BasePatientsListener struct{}

var _ PatientsListener = &BasePatientsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePatientsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePatientsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePatientsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePatientsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePatientsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePatientsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPatients is called when production patients is entered.
func (s *BasePatientsListener) EnterPatients(ctx *PatientsContext) {}

// ExitPatients is called when production patients is exited.
func (s *BasePatientsListener) ExitPatients(ctx *PatientsContext) {}
