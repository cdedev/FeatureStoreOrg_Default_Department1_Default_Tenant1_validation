// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675068711336/Formula.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Formula

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFormulaListener is a complete listener for a parse tree produced by FormulaParser.
type BaseFormulaListener struct{}

var _ FormulaListener = &BaseFormulaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFormulaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFormulaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFormulaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFormulaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFormulaListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFormulaListener) ExitExpression(ctx *ExpressionContext) {}

// EnterFormula is called when production formula is entered.
func (s *BaseFormulaListener) EnterFormula(ctx *FormulaContext) {}

// ExitFormula is called when production formula is exited.
func (s *BaseFormulaListener) ExitFormula(ctx *FormulaContext) {}
