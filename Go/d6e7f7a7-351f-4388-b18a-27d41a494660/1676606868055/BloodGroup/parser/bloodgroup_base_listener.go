// Code generated from /usr/local/lib/Go/d6e7f7a7-351f-4388-b18a-27d41a494660/1676606868055/BloodGroup.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // BloodGroup

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBloodGroupListener is a complete listener for a parse tree produced by BloodGroupParser.
type BaseBloodGroupListener struct{}

var _ BloodGroupListener = &BaseBloodGroupListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBloodGroupListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBloodGroupListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBloodGroupListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBloodGroupListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseBloodGroupListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseBloodGroupListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBloodgroup is called when production bloodgroup is entered.
func (s *BaseBloodGroupListener) EnterBloodgroup(ctx *BloodgroupContext) {}

// ExitBloodgroup is called when production bloodgroup is exited.
func (s *BaseBloodGroupListener) ExitBloodgroup(ctx *BloodgroupContext) {}
