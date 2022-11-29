// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669700897117/BMI.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BMI

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBMIListener is a complete listener for a parse tree produced by BMIParser.
type BaseBMIListener struct{}

var _ BMIListener = &BaseBMIListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBMIListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBMIListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBMIListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBMIListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBMIListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBMIListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBmi is called when production bmi is entered.
func (s *BaseBMIListener) EnterBmi(ctx *BmiContext) {}

// ExitBmi is called when production bmi is exited.
func (s *BaseBMIListener) ExitBmi(ctx *BmiContext) {}
