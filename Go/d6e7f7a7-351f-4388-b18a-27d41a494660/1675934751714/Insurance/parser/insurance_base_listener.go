// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1675934751714/Insurance.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Insurance

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseInsuranceListener is a complete listener for a parse tree produced by InsuranceParser.
type BaseInsuranceListener struct{}

var _ InsuranceListener = &BaseInsuranceListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseInsuranceListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseInsuranceListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseInsuranceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseInsuranceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseInsuranceListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseInsuranceListener) ExitExpression(ctx *ExpressionContext) {}

// EnterInsurance is called when production insurance is entered.
func (s *BaseInsuranceListener) EnterInsurance(ctx *InsuranceContext) {}

// ExitInsurance is called when production insurance is exited.
func (s *BaseInsuranceListener) ExitInsurance(ctx *InsuranceContext) {}
