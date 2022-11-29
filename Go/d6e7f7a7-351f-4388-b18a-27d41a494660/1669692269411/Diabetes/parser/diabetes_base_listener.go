// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669692269411/Diabetes.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Diabetes

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDiabetesListener is a complete listener for a parse tree produced by DiabetesParser.
type BaseDiabetesListener struct{}

var _ DiabetesListener = &BaseDiabetesListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDiabetesListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDiabetesListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDiabetesListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDiabetesListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseDiabetesListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseDiabetesListener) ExitExpression(ctx *ExpressionContext) {}

// EnterDiabetes is called when production diabetes is entered.
func (s *BaseDiabetesListener) EnterDiabetes(ctx *DiabetesContext) {}

// ExitDiabetes is called when production diabetes is exited.
func (s *BaseDiabetesListener) ExitDiabetes(ctx *DiabetesContext) {}
