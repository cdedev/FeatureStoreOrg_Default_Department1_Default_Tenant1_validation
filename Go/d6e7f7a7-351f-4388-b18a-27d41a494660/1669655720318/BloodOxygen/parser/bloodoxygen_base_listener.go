// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1669655720318/BloodOxygen.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodOxygen

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBloodOxygenListener is a complete listener for a parse tree produced by BloodOxygenParser.
type BaseBloodOxygenListener struct{}

var _ BloodOxygenListener = &BaseBloodOxygenListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBloodOxygenListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBloodOxygenListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBloodOxygenListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBloodOxygenListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBloodOxygenListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBloodOxygenListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBloodoxygen is called when production bloodoxygen is entered.
func (s *BaseBloodOxygenListener) EnterBloodoxygen(ctx *BloodoxygenContext) {}

// ExitBloodoxygen is called when production bloodoxygen is exited.
func (s *BaseBloodOxygenListener) ExitBloodoxygen(ctx *BloodoxygenContext) {}
