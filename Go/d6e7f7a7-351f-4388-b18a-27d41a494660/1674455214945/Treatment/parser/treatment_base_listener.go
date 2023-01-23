// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1674455214945/Treatment.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Treatment

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTreatmentListener is a complete listener for a parse tree produced by TreatmentParser.
type BaseTreatmentListener struct{}

var _ TreatmentListener = &BaseTreatmentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTreatmentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTreatmentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTreatmentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTreatmentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseTreatmentListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseTreatmentListener) ExitExpression(ctx *ExpressionContext) {}

// EnterTreatment is called when production treatment is entered.
func (s *BaseTreatmentListener) EnterTreatment(ctx *TreatmentContext) {}

// ExitTreatment is called when production treatment is exited.
func (s *BaseTreatmentListener) ExitTreatment(ctx *TreatmentContext) {}
